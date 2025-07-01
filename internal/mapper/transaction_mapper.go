package mapper

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/constants"
	"bankSystem/internal/model"
)

func TransactionToEntity(transaction *domain.Transaction) *model.TransactionEntity {
	return &model.TransactionEntity{
		Id:        transaction.Id,
		Action:    string(transaction.Action),
		Amount:    transaction.Amount,
		AccountId: transaction.AccountId,
	}
}

func EntityToTransaction(entity *model.TransactionEntity) *domain.Transaction {
	return &domain.Transaction{
		Id:        entity.Id,
		Action:    constants.TransactionType(entity.Action),
		Amount:    entity.Amount,
		AccountId: entity.AccountId,
	}
}
