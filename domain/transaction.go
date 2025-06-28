package domain

import (
	"bankSystem/domain/enums"
)

type Transaction struct {
	Id        string
	Action    enums.TransactionType
	Amount    float64
	AccountId string
}
