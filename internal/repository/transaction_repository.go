package repository

import (
	"bankSystem/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(db *gorm.DB, tx *model.TransactionEntity) error
	GetTransactionsByAccountId(db *gorm.DB, accountId string) ([]model.TransactionEntity, error)
}

type PostgresTransactionRepository struct{}

func NewPostgresTransactionRepository() *PostgresTransactionRepository {
	return &PostgresTransactionRepository{}
}

func (r *PostgresTransactionRepository) SaveTransaction(db *gorm.DB, tx *model.TransactionEntity) error {
	return db.Create(tx).Error
}

func (r *PostgresTransactionRepository) GetTransactionsByAccountId(db *gorm.DB, accountId string) ([]model.TransactionEntity, error) {
	var transactions []model.TransactionEntity
	err := db.Where("account_id = ?", accountId).Find(&transactions).Error
	return transactions, err
}
