package model

type AccountEntity struct {
	Id           string `gorm:"primaryKey"`
	Balance      float64
	Login        string              `gorm:"not null"`
	Transactions []TransactionEntity `gorm:"foreignKey:AccountId;constraint:OnDelete:CASCADE"`
}

func (AccountEntity) TableName() string {
	return "accounts"
}
