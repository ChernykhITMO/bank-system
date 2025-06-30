package service

import (
	domain2 "bankSystem/domain"
	enums2 "bankSystem/domain/enums"
	"bankSystem/mapper"
	interfaces2 "bankSystem/repostitory/interfaces"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository   interfaces2.UserRepository
	friendRepository interfaces2.FriendRepository
}

func NewUserService(userRepo interfaces2.UserRepository, friendRepo interfaces2.FriendRepository) *UserService {
	return &UserService{
		userRepository:   userRepo,
		friendRepository: friendRepo,
	}
}

func (s *UserService) NewUser(login string, name string, sex enums2.Sex, hairColor enums2.Color) (*domain2.User, error) {
	u, err := s.userRepository.GetUser(login)

	if err == nil {
		user := mapper.EntityToUser(u)
		return user, fmt.Errorf("such a user already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user := &domain2.User{
		Login:     login,
		Name:      name,
		Sex:       sex,
		Friends:   []string{},
		HairColor: hairColor,
		Accounts:  []domain2.Account{},
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

func (s *UserService) GetUser(login string) (*domain2.User, error) {
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
