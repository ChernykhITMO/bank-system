package repository

import (
	"bankSystem/internal/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

type FriendRepository interface {
	AddFriends(db *gorm.DB, user, friend string) error
	RemoveFriend(db *gorm.DB, user, friend string) error
	AreFriends(db *gorm.DB, user, friend string) (bool, error)
	GetFriends(db *gorm.DB, login string) ([]string, error)
}

type PostgresFriendRepository struct{}

func NewPostgresFriendRepository() FriendRepository {
	return &PostgresFriendRepository{}
}

func (r *PostgresFriendRepository) AddFriends(db *gorm.DB, user, friend string) error {
	return db.Create(&model.FriendsEntity{
		UserLogin:   user,
		FriendLogin: friend,
		CreatedAt:   time.Now(),
	}).Error
}

func (r *PostgresFriendRepository) RemoveFriend(db *gorm.DB, user, friend string) error {
	return db.Delete(&model.FriendsEntity{}, "user_login = ? AND friends_login = ?", user, friend).Error
}

func (r *PostgresFriendRepository) AreFriends(db *gorm.DB, user, friend string) (bool, error) {
	var f model.FriendsEntity
	err := db.First(&f, "user_login = ? AND friend_login = ?", user, friend).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *PostgresFriendRepository) GetFriends(db *gorm.DB, login string) ([]string, error) {
	var rows []model.FriendsEntity
	err := db.Where("user_login = ?", login).Find(&rows).Error
	if err != nil {
		return nil, err
	}

	friends := make([]string, len(rows))
	for i, row := range rows {
		friends[i] = row.FriendLogin
	}

	return friends, nil
}
