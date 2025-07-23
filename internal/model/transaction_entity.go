package model

type TransactionEntity struct {
	Id            string `gorm:"primaryKey"`
	Action        string
	Amount        float64
	AccountId     string        `gorm:"not null"`
	AccountEntity AccountEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AccountId;references:Id"`
}

func (TransactionEntity) TableName() string {
	return "transactions"
}
