package mapper

import (
	domain2 "bankSystem/domain"
	model2 "bankSystem/model"
)

func AccountToEntity(account *domain2.Account) *model2.AccountEntity {
	return &model2.AccountEntity{
		Id:      account.Id,
		Balance: account.Balance,
		Login:   account.Login,
	}
}

func EntityToAccount(entity *model2.AccountEntity) *domain2.Account {

	return &domain2.Account{
		Id:      entity.Id,
		Balance: entity.Balance,
		Login:   entity.Login,
	}
}
