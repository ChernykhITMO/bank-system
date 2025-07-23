package repository

import (
	"bankSystem/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(db *gorm.DB, login string) (*model.UserEntity, error)
	SaveUser(db *gorm.DB, user *model.UserEntity) error
	DeleteUser(db *gorm.DB, user *model.UserEntity) error
}

type PostgresUserRepository struct{}

func NewPostgresUserRepository() UserRepository {
	return &PostgresUserRepository{}
}

func (r *PostgresUserRepository) GetUser(db *gorm.DB, login string) (*model.UserEntity, error) {
	var user model.UserEntity
	if err := db.
		Preload("Friends").
		Preload("Accounts").
		Where("login = ?", login).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) SaveUser(db *gorm.DB, user *model.UserEntity) error {
	return db.Save(&user).Error
}

func (r *PostgresUserRepository) DeleteUser(db *gorm.DB, user *model.UserEntity) error {
	return db.Delete(&user).Error
}
