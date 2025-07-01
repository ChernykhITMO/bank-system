package domain

import (
	"bankSystem/internal/domain/constants"
)

type Transaction struct {
	Id        string
	Action    constants.TransactionType
	Amount    float64
	AccountId string
}
