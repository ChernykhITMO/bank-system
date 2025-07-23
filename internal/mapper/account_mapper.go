package mapper

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/model"
)

func AccountToEntity(account *domain.Account) *model.AccountEntity {
	return &model.AccountEntity{
		Id:      account.Id,
		Balance: account.Balance,
		Login:   account.Login,
	}
}

func EntityToAccount(entity *model.AccountEntity) *domain.Account {

	return &domain.Account{
		Id:      entity.Id,
		Balance: entity.Balance,
		Login:   entity.Login,
	}
}
