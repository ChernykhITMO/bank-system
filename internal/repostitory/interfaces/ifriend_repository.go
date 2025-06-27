package interfaces

type FriendRepository interface {
	AddFriends(user, friend string) error
	RemoveFriend(user, friend string) error
	AreFriends(user, friend string) (bool, error)
	GetFriends(login string) ([]string, error)
}
