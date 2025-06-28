package mapper

import (
	"bankSystem/domain"
	"bankSystem/domain/enums"
	"bankSystem/model"
)

func TransactionToEntity(transaction *domain.Transaction) *model.TransactionEntity {
	return &model.TransactionEntity{
		Id:     transaction.Id,
		Action: string(transaction.Action),
		Amount: transaction.Amount,
	}
}

func EntityToTransaction(entity *model.TransactionEntity) *domain.Transaction {
	return &domain.Transaction{
		Id:     entity.Id,
		Action: enums.TransactionType(entity.Action),
		Amount: entity.Amount,
	}
}
