package interfaces

import (
	"bankSystem/model"
)

type AccountRepository interface {
	GetAccount(id string) (*model.AccountEntity, error)
	SaveAccount(account *model.AccountEntity) error
}
