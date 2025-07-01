package mapper

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/constants"
	"bankSystem/internal/model"
)

func UserToEntity(user *domain.User) *model.UserEntity {
	friendEntities := StringToFriendsEntity(user.Login, user.Friends)

	accounts := make([]model.AccountEntity, 0, len(user.Accounts))
	for _, acc := range user.Accounts {
		account := AccountToEntity(&acc)
		accounts = append(accounts, *account)
	}

	return &model.UserEntity{
		Login:     user.Login,
		Name:      user.Name,
		Sex:       string(user.Sex),
		HairColor: string(user.HairColor),
		Friends:   friendEntities,
		Accounts:  accounts,
	}
}

func EntityToUser(entity *model.UserEntity) *domain.User {
	friendLogins := make([]string, 0, len(entity.Friends))
	for _, friend := range entity.Friends {
		friendLogins = append(friendLogins, friend.FriendLogin)
	}

	accounts := make([]domain.Account, 0, len(entity.Accounts))
	for _, acc := range entity.Accounts {
		account := EntityToAccount(&acc)
		accounts = append(accounts, *account)
	}

	return &domain.User{
		Login:     entity.Login,
		Name:      entity.Name,
		Sex:       constants.Sex(entity.Sex),
		HairColor: constants.Color(entity.HairColor),
		Friends:   friendLogins,
		Accounts:  accounts,
	}
}
