package model

type AccountEntity struct {
	Id      string `gorm:"primaryKey"`
	Balance float64
	Login   string
	History []TransactionEntity `gorm:"foreignKey:AccountId;references:Id"`
}

func (AccountEntity) TableName() string {
	return "accounts"
}
