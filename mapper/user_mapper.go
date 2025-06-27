package mapper

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/domain/enums"
	"bankSystem/internal/model"
)

func UserToEntity(user *domain.User) *model.UserEntity {
	return &model.UserEntity{
		Login:     user.Login,
		Name:      user.Name,
		Sex:       string(user.Sex),
		HairColor: string(user.HairColor),
	}
}

func EntityToUser(entity *model.UserEntity) *domain.User {
	return &domain.User{
		Login:     entity.Login,
		Name:      entity.Name,
		Sex:       enums.Sex(entity.Sex),
		HairColor: enums.Color(entity.HairColor),
	}
}
