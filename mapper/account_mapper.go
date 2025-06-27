package mapper

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/model"
)

func AccountToEntity(account *domain.Account) *model.AccountEntity {
	history := make([]model.TransactionEntity, len(account.History))
	for i, tx := range account.History {
		history[i] = *TransactionToEntity(&tx)
	}

	return &model.AccountEntity{
		Id:      account.Id,
		Balance: account.Balance,
		Login:   account.Login,
		History: history,
	}
}

func EntityToAccount(entity *model.AccountEntity) *domain.Account {
	history := make([]domain.Transaction, len(entity.History))
	for i, tx := range entity.History {
		history[i] = *EntityToTransaction(&tx)
	}

	return &domain.Account{
		Id:      entity.Id,
		Balance: entity.Balance,
		Login:   entity.Login,
		History: history,
	}
}
