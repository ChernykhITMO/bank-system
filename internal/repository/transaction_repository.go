package repository

import (
	"bankSystem/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(tx *model.TransactionEntity) error
	GetTransactionsByAccountId(accountId string) ([]model.TransactionEntity, error)
}

type PostgresTransactionRepository struct {
	db *gorm.DB
}

func NewPostgresTransactionRepository(db *gorm.DB) *PostgresTransactionRepository {
	return &PostgresTransactionRepository{db: db}
}

func (r *PostgresTransactionRepository) SaveTransaction(tx *model.TransactionEntity) error {
	return r.db.Create(tx).Error
}

func (r *PostgresTransactionRepository) GetTransactionsByAccountId(accountId string) ([]model.TransactionEntity, error) {
	var transactions []model.TransactionEntity
	err := r.db.Where("account_id = ?", accountId).Find(&transactions).Error
	return transactions, err
}
