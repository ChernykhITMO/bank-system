package interfaces

import (
	"bankSystem/model"
)

type UserRepository interface {
	GetUser(login string) (*model.UserEntity, error)
	SaveUser(user *model.UserEntity) error
	DeleteUser(user *model.UserEntity) error
}
