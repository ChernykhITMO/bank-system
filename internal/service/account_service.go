package service

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/constants"
	"bankSystem/internal/mapper"
	"bankSystem/internal/repository"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountService struct {
	tx               repository.TxManager
	repoAccount      repository.AccountRepository
	repoUser         repository.UserRepository
	repoFriends      repository.FriendRepository
	repoTransactions repository.TransactionRepository
}

func NewAccountService(tx repository.TxManager, repoAcc repository.AccountRepository, repoUser repository.UserRepository,
	repoFriends repository.FriendRepository, repoTransactions repository.TransactionRepository) *AccountService {
	return &AccountService{
		tx:               tx,
		repoAccount:      repoAcc,
		repoUser:         repoUser,
		repoFriends:      repoFriends,
		repoTransactions: repoTransactions,
	}
}

func (s *AccountService) NewUserAccount(user *domain.User) error {
	return s.tx.WithTx(func(tx *gorm.DB) error {
		_, ok := s.repoUser.GetUser(tx, user.Login)
		if ok != nil {
			return fmt.Errorf("User not found")
		}

		account := domain.Account{
			Id:    uuid.NewString(),
			Login: user.Login,
		}

		user.Accounts = append(user.Accounts, account)
		accountEntity := mapper.AccountToEntity(&account)
		return s.repoAccount.SaveAccount(tx, accountEntity)
	})
}

func (s *AccountService) GetBalance(id string) (float64, error) {
	var balance float64

	err := s.tx.WithTx(func(tx *gorm.DB) error {
		account, err := s.repoAccount.GetAccount(tx, id)
		if err != nil {
			return err
		}

		balance = account.Balance
		return nil
	})

	return balance, err
}

func (s *AccountService) Deposit(accountID string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	return s.tx.WithTx(func(tx *gorm.DB) error {
		accountEntity, err := s.repoAccount.GetAccount(tx, accountID)
		if err != nil {
			return fmt.Errorf("account not found")
		}

		account := mapper.EntityToAccount(accountEntity)
		account.Balance += amount

		if err = s.repoAccount.SaveAccount(tx, mapper.AccountToEntity(account)); err != nil {
			return err
		}

		transaction := domain.Transaction{
			Id:              uuid.NewString(),
			TransactionType: constants.TransactionDeposit,
			Amount:          amount,
			AccountId:       account.Id,
		}

		return s.repoTransactions.SaveTransaction(tx, mapper.TransactionToEntity(&transaction))
	})
}

func (s *AccountService) Withdraw(id string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Amount can't be negative")
	}

	return s.tx.WithTx(func(tx *gorm.DB) error {
		accountEntity, err := s.repoAccount.GetAccount(tx, id)
		if err != nil {
			return fmt.Errorf("Account not found")
		}
		if accountEntity.Balance < amount {
			return fmt.Errorf("Account doesn't have this amount")
		}

		account := mapper.EntityToAccount(accountEntity)
		account.Balance -= amount

		if err := s.repoAccount.SaveAccount(tx, mapper.AccountToEntity(account)); err != nil {
			return err
		}

		transaction := domain.Transaction{
			Id:              uuid.NewString(),
			TransactionType: constants.TransactionWithdraw,
			Amount:          amount,
			AccountId:       account.Id,
		}

		return s.repoTransactions.SaveTransaction(tx, mapper.TransactionToEntity(&transaction))
	})
}

func (s *AccountService) Transfer(id1, id2 string, amount float64) error {
	return s.tx.WithTx(func(tx *gorm.DB) error {
		accountEntity1, err := s.repoAccount.GetAccount(tx, id1)
		if err != nil {
			return fmt.Errorf("Source account not found")
		}

		accountEntity2, err := s.repoAccount.GetAccount(tx, id2)
		if err != nil {
			return fmt.Errorf("Target account not found")
		}

		user1 := accountEntity1.Login
		user2 := accountEntity2.Login

		isFriends, err := s.repoFriends.AreFriends(tx, user1, user2)
		if err != nil {
			return err
		}

		account1 := mapper.EntityToAccount(accountEntity1)
		account2 := mapper.EntityToAccount(accountEntity2)

		var fee float64
		if isFriends {
			fee = amount * 0.03
		} else if user1 != user2 {
			fee = amount * 0.10
		}

		total := amount + fee
		if account1.Balance < total {
			return fmt.Errorf("Insufficient funds")
		}

		account1.Balance -= total
		account2.Balance += amount

		if err := s.repoAccount.SaveAccount(tx, mapper.AccountToEntity(account1)); err != nil {
			return err
		}
		if err := s.repoAccount.SaveAccount(tx, mapper.AccountToEntity(account2)); err != nil {
			return err
		}

		outgoing := domain.Transaction{
			Id:              uuid.NewString(),
			TransactionType: constants.TransactionTransfer,
			Amount:          -total,
			AccountId:       account1.Id,
		}

		incoming := domain.Transaction{
			Id:              uuid.NewString(),
			TransactionType: constants.TransactionTransfer,
			Amount:          amount,
			AccountId:       account2.Id,
		}

		if err := s.repoTransactions.SaveTransaction(tx, mapper.TransactionToEntity(&outgoing)); err != nil {
			return err
		}
		return s.repoTransactions.SaveTransaction(tx, mapper.TransactionToEntity(&incoming))
	})
}

func (s *AccountService) DeleteAccount(id, login string) error {
	return s.tx.WithTx(func(tx *gorm.DB) error {
		_, err := s.repoUser.GetUser(tx, login)
		if err != nil {
			return fmt.Errorf("User not found")
		}

		_, err = s.repoAccount.GetAccount(tx, id)
		if err != nil {
			return fmt.Errorf("Account not found")
		}

		return s.repoAccount.DeleteAccount(tx, id)
	})
}

func (s *AccountService) GetTransactions(id string) (*[]domain.Transaction, error) {
	var result []domain.Transaction

	err := s.tx.WithTx(func(tx *gorm.DB) error {
		_, err := s.repoAccount.GetAccount(tx, id)
		if err != nil {
			return err
		}

		transactionEntities, err := s.repoTransactions.GetTransactionsByAccountId(tx, id)
		if err != nil {
			return err
		}

		for _, te := range transactionEntities {
			result = append(result, *mapper.EntityToTransaction(&te))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &result, nil
}
