package repository

import "gorm.io/gorm"

type TxManager interface {
	WithTx(fn func(tx *gorm.DB) error) error
}

type GormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(db *gorm.DB) TxManager {
	return &GormTxManager{db: db}
}

func (t *GormTxManager) WithTx(fn func(tx *gorm.DB) error) error {
	return t.db.Transaction(fn)
}
