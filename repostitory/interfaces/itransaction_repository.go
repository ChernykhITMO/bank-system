package interfaces

import "bankSystem/model"

type TransactionRepository interface {
	SaveTransaction(tx *model.TransactionEntity) error
	GetTransactionsByAccountId(accountId string) ([]model.TransactionEntity, error)
}
