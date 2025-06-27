package service

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/enums"
	"bankSystem/internal/repostitory/interfaces"
	"bankSystem/mapper"
	"fmt"
	"github.com/google/uuid"
)

type AccountService struct {
	repoAccount interfaces.AccountRepository
	repoUser    interfaces.UserRepository
	repoFriends interfaces.FriendRepository
}

func NewAccountService(repoAcc interfaces.AccountRepository, repoUser interfaces.UserRepository, repoFriends interfaces.FriendRepository) *AccountService {
	return &AccountService{
		repoAccount: repoAcc,
		repoUser:    repoUser,
		repoFriends: repoFriends,
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
	s.repoAccount.SaveAccount(accountEntity)
	fmt.Println("Created account with Id:", account.Id)
	return nil
}

func (s *AccountService) GetBalance(account *domain.Account) float64 {
	return account.Balance
}

func (s *AccountService) Deposit(account *domain.Account, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Amount can't be negavtive")
	}

	transaction := domain.Transaction{
		Id:        uuid.NewString(),
		Action:    enums.TransactionDeposit,
		Amount:    amount,
		AccountId: account.Id,
	}

	account.History = append(account.History, transaction)
	account.Balance += amount
	s.repoAccount.SaveAccount(mapper.AccountToEntity(account))
	return nil
}

func (s *AccountService) Withdraw(account *domain.Account, amount float64) error {
	if account.Balance < amount {
		return fmt.Errorf("Account doesn't have this amount")
	}
	if amount <= 0 {
		return fmt.Errorf("Amount can't be negative")
	}

	transaction := domain.Transaction{
		Id:        uuid.NewString(),
		Action:    enums.TransactionWithdraw,
		Amount:    amount,
		AccountId: account.Id,
	}

	account.History = append(account.History, transaction)

	account.Balance -= amount
	s.repoAccount.SaveAccount(mapper.AccountToEntity(account))
	return nil
}

func (s *AccountService) Transfer(account1, account2 *domain.Account, amount float64) error {
	user1Login, user2Login := account1.Login, account2.Login

	isFriends, err := s.repoFriends.AreFriends(user1Login, user2Login)
	if err != nil {
		return err
	}
	if isFriends {
		if account1.Balance < amount+amount*0.03 {
			return fmt.Errorf("account %s doesn't have amount", user1Login)
		}
		account1.Balance -= (amount + amount*0.03)
		account2.Balance += amount

		account1.History = append(account1.History, domain.Transaction{
			Id:     uuid.NewString(),
			Action: enums.TransactionTransfer,
			Amount: -(amount + amount*0.03),
		})

	} else if user1Login == user2Login {
		if account1.Balance < amount {
			return fmt.Errorf("account %s doesn't have amount", user1Login)
		}
		account1.Balance -= amount
		account2.Balance += amount

		account1.History = append(account1.History, domain.Transaction{
			Id:     uuid.NewString(),
			Action: enums.TransactionTransfer,
			Amount: -amount,
		})
	} else {
		if account1.Balance < amount+amount*0.10 {
			return fmt.Errorf("account %s doesn't have amount", user1Login)
		}
		account1.Balance -= amount + amount*0.10
		account2.Balance += amount

		account1.History = append(account1.History, domain.Transaction{
			Id:     uuid.NewString(),
			Action: enums.TransactionTransfer,
			Amount: -(amount + amount*0.10),
		})
	}
	account2.History = append(account2.History, domain.Transaction{
		Id:     uuid.NewString(),
		Action: enums.TransactionTransfer,
		Amount: amount,
	})
	return nil
}
