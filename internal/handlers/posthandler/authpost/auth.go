package authpost

import (
	"io"
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

type UserModel interface {
	Delete(id string) error
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
	Insert(user data.User) error
	Update(id string, info data.UpdateUser) error
	EmailExists(email string) (bool, error)
	IDExists(id string) (bool, error)
}

type PreferenceModel interface {
	Exists(userId string) (bool, error)
	Get(userId string) (data.Preference, error)
	Insert(preference data.Preference) error
	Update(preference data.PreferenceUpdate) error
}

type StorageModel interface {
	DeleteProfilePicture(relativeFilePath ...string) error
	GetPublicUrl(relativeFilePath string) string
	UpdateProfilePicture(relativeFilePath string, file io.Reader) error
	UploadProfilePicture(fileRelativePath string, file io.Reader) error
	Exists(folder string, target string) (bool, error)
}

type Auth struct {
	logger          jsonlog.Logger
	usermodel       UserModel
	preferenceModel PreferenceModel
	embedded        groupietracker.Embedded
	storage         StorageModel
}

func New(logger jsonlog.Logger, userModel UserModel, preferenceModel PreferenceModel, storage StorageModel, embedded groupietracker.Embedded) *Auth {
	return &Auth{
		logger:          logger,
		usermodel:       userModel,
		preferenceModel: preferenceModel,
		storage:         storage,
		embedded:        embedded,
	}
}

func (a *Auth) getUser(r *http.Request) data.User {
	val := r.Context().Value(utils.USER_ID_KEY)

	if user, ok := val.(data.User); ok {
		return user
	}
	return data.User{}
}
