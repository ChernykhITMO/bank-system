package mapper

import (
	domain2 "bankSystem/domain"
	model2 "bankSystem/model"
)

func AccountToEntity(account *domain2.Account) *model2.AccountEntity {
	history := make([]model2.TransactionEntity, len(account.History))
	for i, tx := range account.History {
		history[i] = *TransactionToEntity(&tx)
	}

	return &model2.AccountEntity{
		Id:      account.Id,
		Balance: account.Balance,
		Login:   account.Login,
		History: history,
	}
}

func EntityToAccount(entity *model2.AccountEntity) *domain2.Account {
	history := make([]domain2.Transaction, len(entity.History))
	for i, tx := range entity.History {
		history[i] = *EntityToTransaction(&tx)
	}

	return &domain2.Account{
		Id:      entity.Id,
		Balance: entity.Balance,
		Login:   entity.Login,
		History: history,
	}
}
