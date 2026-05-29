package posthandler

import "github.com/dositadi/groupie-tracker/internal/data"

type UserModel interface {
	Delete(id string) error
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
	Insert(user data.User) error
	Update(id string, info data.UpdateUser) error
	EmailExists(email string) (bool, error)
	IDExists(id string) (bool, error)
}

type Post struct {
	userModel UserModel
}

func New(userModel UserModel) *Post {
	return &Post{
		userModel: userModel,
	}
}
