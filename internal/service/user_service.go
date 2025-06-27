package service

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/enums"
	"bankSystem/internal/repostitory/interfaces"
	"bankSystem/mapper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository   interfaces.UserRepository
	friendRepository interfaces.FriendRepository
}

func NewUserService(userRepo interfaces.UserRepository, friendRepo interfaces.FriendRepository) *UserService {
	return &UserService{
		userRepository:   userRepo,
		friendRepository: friendRepo,
	}
}

func (s *UserService) NewUser(login string, name string, sex enums.Sex, hairColor enums.Color) (*domain.User, error) {
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
	s.userRepository.SaveUser(entity)

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

func contains(friends []string, login string) bool {
	for _, user := range friends {
		if user == login {
			return true
		}
	}
	return false
}

func (s *UserService) GetUser(login string) (*domain.User, error) {
	userEntity, err := s.userRepository.GetUser(login)
	if err != nil {
		return nil, err
	}

	user := mapper.EntityToUser(userEntity)

	friends, err := s.friendRepository.GetFriends(login)
	if err != nil {
		return nil, err
	}

	user.Friends = friends

	return user, nil
}
