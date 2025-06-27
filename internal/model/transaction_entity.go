package model

type TransactionEntity struct {
	Id        string `gorm:"primaryKey"`
	Action    string
	Amount    float64
	AccountId string
}

func (TransactionEntity) TableName() string {
	return "transactions"
}
