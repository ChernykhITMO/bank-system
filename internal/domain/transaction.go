package domain

import (
	"bankSystem/internal/domain/enums"
)

type Transaction struct {
	Id        string
	Action    enums.TransactionType
	Amount    float64
	AccountId string
}
