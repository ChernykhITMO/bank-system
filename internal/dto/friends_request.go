package dto

type FriendsRequest struct {
	UserLogin   string `json:"user_login" binding:"required"`
	FriendLogin string `json:"friend_login" binding:"required"`
}
