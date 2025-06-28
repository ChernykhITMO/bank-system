package repostitory

import (
	"bankSystem/model"
	"gorm.io/gorm"
)

type PostgresAccountRepository struct {
	db *gorm.DB
}

func NewPostgresAccountRepository(db *gorm.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) GetAccount(id string) (*model.AccountEntity, error) {
	var account model.AccountEntity
	if err := r.db.Find("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *PostgresAccountRepository) SaveAccount(account *model.AccountEntity) error {
	return r.db.Create(account).Error
}
