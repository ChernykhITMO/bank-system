package repository

import (
	"bankSystem/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(login string) (*model.UserEntity, error)
	SaveUser(user *model.UserEntity) error
	DeleteUser(user *model.UserEntity) error
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(dataBase *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: dataBase,
	}
}

func (r *PostgresUserRepository) GetUser(login string) (*model.UserEntity, error) {
	var user model.UserEntity
	if err := r.db.
		Preload("Friends").
		Preload("Accounts").
		Where("login = ?", login).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) SaveUser(user *model.UserEntity) error {
	return r.db.Save(&user).Error
}

func (r *PostgresUserRepository) DeleteUser(user *model.UserEntity) error {
	return r.db.Delete(&user).Error
}
