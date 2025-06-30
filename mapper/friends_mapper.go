package mapper

import "bankSystem/model"

func StringToFriendsEntity(userLogin string, friends []string) []model.FriendsEntity {
	friendEntities := make([]model.FriendsEntity, 0, len(friends))
	for _, friendLogin := range friends {
		friendEntities = append(friendEntities, model.FriendsEntity{
			UserLogin:   userLogin,
			FriendLogin: friendLogin,
		})
	}
	return friendEntities
}
