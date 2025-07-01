package repository

import (
	"bankSystem/internal/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

type FriendRepository interface {
	AddFriends(user, friend string) error
	RemoveFriend(user, friend string) error
	AreFriends(user, friend string) (bool, error)
	GetFriends(login string) ([]string, error)
}

type PostgresFriendRepository struct {
	db *gorm.DB
}

func NewPostgresFriendRepository(db *gorm.DB) *PostgresFriendRepository {
	return &PostgresFriendRepository{db: db}
}

func (r *PostgresFriendRepository) AddFriends(user, friend string) error {
	return r.db.Create(&model.FriendsEntity{
		UserLogin:   user,
		FriendLogin: friend,
		CreatedAt:   time.Now(),
	}).Error
}

func (r *PostgresFriendRepository) RemoveFriend(user, friend string) error {
	return r.db.Delete(&model.FriendsEntity{}, "user_login = ? AND friends_login = ?", user, friend).Error
}

func (r *PostgresFriendRepository) AreFriends(user, friend string) (bool, error) {
	var f model.FriendsEntity
	err := r.db.First(&f, "user_login = ? AND friend_login = ?", user, friend).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *PostgresFriendRepository) GetFriends(login string) ([]string, error) {
	var rows []model.FriendsEntity
	err := r.db.Where("user_login = ?", login).Find(&rows).Error
	if err != nil {
		return nil, err
	}

	friends := make([]string, len(rows))
	for i, row := range rows {
		friends[i] = row.FriendLogin
	}

	return friends, nil
}
