package model

type UserEntity struct {
	Login     string `gorm:"primaryKey"`
	Name      string
	Sex       string
	HairColor string
	Friends   []FriendsEntity `gorm:"foreignKey:UserLogin;references:Login"`
	Accounts  []AccountEntity `gorm:"foreignKey:Login;references:Login"`
}

func (UserEntity) TableName() string {
	return "users"
}
