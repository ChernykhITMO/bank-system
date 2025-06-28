package mapper

import (
	"bankSystem/domain"
	enums2 "bankSystem/domain/enums"
	"bankSystem/model"
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
		Sex:       enums2.Sex(entity.Sex),
		HairColor: enums2.Color(entity.HairColor),
	}
}
