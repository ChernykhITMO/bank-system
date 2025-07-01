package service

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/constants"
	"bankSystem/internal/mapper"
	repository2 "bankSystem/internal/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository   repository2.UserRepository
	friendRepository repository2.FriendRepository
}

func NewUserService(userRepo repository2.UserRepository, friendRepo repository2.FriendRepository) *UserService {
	return &UserService{
		userRepository:   userRepo,
		friendRepository: friendRepo,
	}
}

func (s *UserService) NewUser(login string, name string, sex constants.Sex, hairColor constants.Color) (*domain.User, error) {
	u, err := s.userRepository.GetUser(login)

	if err == nil {
		user := mapper.EntityToUser(u)
		return user, fmt.Errorf("such a user already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user := &domain.User{
		Login:     login,
		Name:      name,
		Sex:       sex,
		Friends:   []string{},
		HairColor: hairColor,
		Accounts:  []domain.Account{},
	}

	entity := mapper.UserToEntity(user)
	err = s.userRepository.SaveUser(entity)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) AddFriend(user, friend string) error {
	already, err := s.friendRepository.AreFriends(user, friend)

	if err != nil {
		return err
	}

	if already {
		fmt.Errorf("Users are already friends.")
	}

	if err := s.friendRepository.AddFriends(user, friend); err != nil {
		return err
	}

	return s.friendRepository.AddFriends(friend, user)
}

func (s *UserService) RemoveFriend(userLogin, friendLogin string) error {
	return s.friendRepository.RemoveFriend(friendLogin, userLogin)
}

func (s *UserService) GetUser(login string) (*domain.User, error) {
	userEntity, err := s.userRepository.GetUser(login)
	if err != nil {
		return nil, err
	}

	user := mapper.EntityToUser(userEntity)

	return user, nil
}

func (s *UserService) DeleteUser(login string) error {
	userEntity, err := s.userRepository.GetUser(login)
	if err != nil {
		return err
	}
	ok := s.userRepository.DeleteUser(userEntity)
	if ok != nil {
		return ok
	}
	return nil
}
