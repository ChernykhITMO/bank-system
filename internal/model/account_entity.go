package model

type AccountEntity struct {
	Id      string `gorm:"primaryKey"`
	Balance float64
	Login   string
}

func (AccountEntity) TableName() string {
	return "accounts"
}
