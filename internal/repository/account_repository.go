package repository

import (
	model2 "bankSystem/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccount(db *gorm.DB, id string) (*model2.AccountEntity, error)
	SaveAccount(db *gorm.DB, account *model2.AccountEntity) error
	DeleteAccount(db *gorm.DB, id string) error
}

type PostgresAccountRepository struct{}

func NewPostgresAccountRepository() AccountRepository {
	return &PostgresAccountRepository{}
}

func (r *PostgresAccountRepository) GetAccount(db *gorm.DB, id string) (*model2.AccountEntity, error) {
	var account model2.AccountEntity
	if err := db.Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *PostgresAccountRepository) SaveAccount(db *gorm.DB, account *model2.AccountEntity) error {
	return db.Save(account).Error
}

func (r *PostgresAccountRepository) DeleteAccount(db *gorm.DB, id string) error {
	return db.Delete(&model2.AccountEntity{}, "id = ?", id).Error
}
