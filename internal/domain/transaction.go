package domain

import (
	"bankSystem/internal/domain/constants"
)

type Transaction struct {
	Id              string
	TransactionType constants.TransactionType
	Amount          float64
	AccountId       string
}
