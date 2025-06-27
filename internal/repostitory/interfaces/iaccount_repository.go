package interfaces

import (
	"bankSystem/internal/model"
)

type AccountRepository interface {
	GetAccount(id string) (*model.AccountEntity, error)
	SaveAccount(account *model.AccountEntity) error
}
