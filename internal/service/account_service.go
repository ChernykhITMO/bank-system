package service

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/constants"
	"bankSystem/internal/mapper"
	"bankSystem/internal/repository"
	"fmt"
	"github.com/google/uuid"
)

type AccountService struct {
	repoAccount      repository.AccountRepository
	repoUser         repository.UserRepository
	repoFriends      repository.FriendRepository
	repoTransactions repository.TransactionRepository
}

func NewAccountService(repoAcc repository.AccountRepository, repoUser repository.UserRepository,
	repoFriends repository.FriendRepository, repoTransactions repository.TransactionRepository) *AccountService {
	return &AccountService{
		repoAccount:      repoAcc,
		repoUser:         repoUser,
		repoFriends:      repoFriends,
		repoTransactions: repoTransactions,
	}
}

func (s *AccountService) NewUserAccount(user *domain.User) error {
	_, ok := s.repoUser.GetUser(user.Login)
	if ok != nil {
		return fmt.Errorf("User not found")
	}

	account := domain.Account{
		Id:    uuid.NewString(),
		Login: user.Login,
	}

	user.Accounts = append(user.Accounts, account)
	accountEntity := mapper.AccountToEntity(&account)
	err := s.repoAccount.SaveAccount(accountEntity)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountService) GetBalance(id string) (float64, error) {
	account, err := s.repoAccount.GetAccount(id)
	if err != nil {
		return 0, fmt.Errorf("Account not found")
	}
	return account.Balance, nil
}

func (s *AccountService) Deposit(id string, amount float64) error {
	accountEntity, err := s.repoAccount.GetAccount(id)
	if err != nil {
		return fmt.Errorf("Account not found")
	}
	if amount <= 0 {
		return fmt.Errorf("Amount can't be negavtive")
	}

	transaction := domain.Transaction{
		Id:        uuid.NewString(),
		Action:    constants.TransactionDeposit,
		Amount:    amount,
		AccountId: accountEntity.Id,
	}

	account := mapper.EntityToAccount(accountEntity)
	account.Balance += amount

	err = s.repoAccount.SaveAccount(mapper.AccountToEntity(account))
	if err != nil {
		return err
	}

	transactionEntity := mapper.TransactionToEntity(&transaction)
	err = s.repoTransactions.SaveTransaction(transactionEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) Withdraw(id string, amount float64) error {
	accountEntity, err := s.repoAccount.GetAccount(id)
	if err != nil {
		return fmt.Errorf("Account not found")
	}

	if accountEntity.Balance < amount {
		return fmt.Errorf("Account doesn't have this amount")
	}
	if amount <= 0 {
		return fmt.Errorf("Amount can't be negative")
	}

	transaction := domain.Transaction{
		Id:        uuid.NewString(),
		Action:    constants.TransactionWithdraw,
		Amount:    amount,
		AccountId: accountEntity.Id,
	}

	account := mapper.EntityToAccount(accountEntity)

	account.Balance -= amount
	err = s.repoAccount.SaveAccount(mapper.AccountToEntity(account))
	if err != nil {
		return err
	}

	transactionEntity := mapper.TransactionToEntity(&transaction)
	err = s.repoTransactions.SaveTransaction(transactionEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) Transfer(id1, id2 string, amount float64) error {
	accountEntity1, err := s.repoAccount.GetAccount(id1)
	if err != nil {
		return fmt.Errorf("Account not found")
	}

	accountEntity2, err := s.repoAccount.GetAccount(id2)
	if err != nil {
		return fmt.Errorf("Account not found")
	}

	user1Login, user2Login := accountEntity1.Login, accountEntity2.Login

	isFriends, err := s.repoFriends.AreFriends(user1Login, user2Login)
	if err != nil {
		return err
	}

	account1 := mapper.EntityToAccount(accountEntity1)
	account2 := mapper.EntityToAccount(accountEntity2)

	var fee float64
	if isFriends {
		fee = amount * 0.03
	} else if user1Login != user2Login {
		fee = amount * 0.10
	}

	total := amount + fee
	if account1.Balance < total {
		return fmt.Errorf("account %s doesn't have enough funds", user1Login)
	}

	account1.Balance -= total
	account2.Balance += amount

	outgoingTx := domain.Transaction{
		Id:        uuid.NewString(),
		Action:    constants.TransactionTransfer,
		Amount:    -total,
		AccountId: account1.Id,
	}

	incomingTx := domain.Transaction{
		Id:        uuid.NewString(),
		Action:    constants.TransactionTransfer,
		Amount:    amount,
		AccountId: account2.Id,
	}

	if err = s.repoAccount.SaveAccount(mapper.AccountToEntity(account1)); err != nil {
		return err
	}
	if err = s.repoAccount.SaveAccount(mapper.AccountToEntity(account2)); err != nil {
		return err
	}

	if err = s.repoTransactions.SaveTransaction(mapper.TransactionToEntity(&outgoingTx)); err != nil {
		return err
	}
	if err = s.repoTransactions.SaveTransaction(mapper.TransactionToEntity(&incomingTx)); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) DeleteAccount(id, login string) error {
	_, ok := s.repoUser.GetUser(login)
	if ok != nil {
		return fmt.Errorf("User not found")
	}

	_, err := s.repoAccount.GetAccount(id)
	if err != nil {
		return fmt.Errorf("Account not found")
	}
	err = s.repoAccount.DeleteAccount(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountService) GetTransactions(id string) (*[]domain.Transaction, error) {
	_, err := s.repoAccount.GetAccount(id)
	if err != nil {
		return nil, fmt.Errorf("Account not found")
	}

	transactionEntities, err := s.repoTransactions.GetTransactionsByAccountId(id)
	if err != nil {
		return nil, err
	}

	var transactions []domain.Transaction
	for _, te := range transactionEntities {
		transactions = append(transactions, *mapper.EntityToTransaction(&te))
	}

	return &transactions, nil
}
