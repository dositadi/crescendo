package deletehandler

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

type Delete struct {
	userModel UserModel
}

func New(userModel UserModel) *Delete {
	return &Delete{
		userModel: userModel,
	}
}
