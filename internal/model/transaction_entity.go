package model

type TransactionEntity struct {
	Id        string `gorm:"primaryKey"`
	Action    string
	Amount    float64
	AccountId string `gorm:"not null"`
}

func (TransactionEntity) TableName() string {
	return "transactions"
}
