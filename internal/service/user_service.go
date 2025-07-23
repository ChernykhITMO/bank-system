package service

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/constants"
	"bankSystem/internal/mapper"
	"bankSystem/internal/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserService struct {
	tx               repository.TxManager
	userRepository   repository.UserRepository
	friendRepository repository.FriendRepository
}

func NewUserService(tx repository.TxManager, userRepo repository.UserRepository, friendRepo repository.FriendRepository) *UserService {
	return &UserService{
		tx:               tx,
		userRepository:   userRepo,
		friendRepository: friendRepo,
	}
}

func (s *UserService) NewUser(login, name string, sex constants.Sex, hairColor constants.Color) (*domain.User, error) {
	var user domain.User

	err := s.tx.WithTx(func(tx *gorm.DB) error {
		_, err := s.userRepository.GetUser(tx, login)
		if err == nil {
			return fmt.Errorf("user with login '%s' already exists", login)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		user = domain.User{
			Login:     login,
			Name:      name,
			Sex:       sex,
			HairColor: hairColor,
			Friends:   []string{},
			Accounts:  []domain.Account{},
		}

		return s.userRepository.SaveUser(tx, mapper.UserToEntity(&user))
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) AddFriend(userLogin, friendLogin string) error {
	return s.tx.WithTx(func(tx *gorm.DB) error {
		already, err := s.friendRepository.AreFriends(tx, userLogin, friendLogin)
		if err != nil {
			return err
		}
		if already {
			return fmt.Errorf("users '%s' and '%s' are already friends", userLogin, friendLogin)
		}

		if err := s.friendRepository.AddFriends(tx, userLogin, friendLogin); err != nil {
			return err
		}
		return s.friendRepository.AddFriends(tx, friendLogin, userLogin)
	})
}

func (s *UserService) RemoveFriend(userLogin, friendLogin string) error {
	return s.tx.WithTx(func(tx *gorm.DB) error {
		if err := s.friendRepository.RemoveFriend(tx, userLogin, friendLogin); err != nil {
			return err
		}
		return s.friendRepository.RemoveFriend(tx, friendLogin, userLogin)
	})
}

func (s *UserService) GetUser(login string) (*domain.User, error) {
	var user *domain.User

	err := s.tx.WithTx(func(tx *gorm.DB) error {
		userEntity, err := s.userRepository.GetUser(tx, login)
		if err != nil {
			return err
		}
		user = mapper.EntityToUser(userEntity)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(login string) error {
	return s.tx.WithTx(func(tx *gorm.DB) error {
		userEntity, err := s.userRepository.GetUser(tx, login)
		if err != nil {
			return err
		}
		return s.userRepository.DeleteUser(tx, userEntity)
	})
}
