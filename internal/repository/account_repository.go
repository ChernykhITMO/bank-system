package repository

import (
	model2 "bankSystem/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccount(id string) (*model2.AccountEntity, error)
	SaveAccount(account *model2.AccountEntity) error
	DeleteAccount(id string) error
}

type PostgresAccountRepository struct {
	db *gorm.DB
}

func NewPostgresAccountRepository(db *gorm.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) GetAccount(id string) (*model2.AccountEntity, error) {
	var account model2.AccountEntity
	if err := r.db.Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *PostgresAccountRepository) SaveAccount(account *model2.AccountEntity) error {
	return r.db.Save(account).Error
}

func (r *PostgresAccountRepository) DeleteAccount(id string) error {
	if err := r.db.Where("account_id = ?", id).Delete(&model2.TransactionEntity{}).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model2.AccountEntity{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
