package model

import "time"

type FriendsEntity struct {
	UserLogin   string `gorm:"primaryKey"`
	FriendLogin string `gorm:"primaryKey"`
	CreatedAt   time.Time
}

func (FriendsEntity) TableName() string {
	return "friends"
}
