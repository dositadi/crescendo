package models

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	favoritemodel "github.com/dositadi/groupie-tracker/internal/models/favorite_model"
	preferencemodel "github.com/dositadi/groupie-tracker/internal/models/preference_model"
	"github.com/dositadi/groupie-tracker/internal/models/soldticketsmodel"
	usermodel "github.com/dositadi/groupie-tracker/internal/models/user_model"
	"github.com/jackc/pgx/v5"
)

type Models struct {
	UserModel        usermodel.UserModel
	FavoriteModel    favoritemodel.FavoriteModel
	SoldTicketsModel soldticketsmodel.SoldTicketsModel
	PreferenceModel  preferencemodel.PreferenceModel
}

func New(db *pgx.Conn, logger jsonlog.Logger) *Models {
	return &Models{
		UserModel:        *usermodel.New(db, logger),
		FavoriteModel:    *favoritemodel.New(db, logger),
		SoldTicketsModel: *soldticketsmodel.New(db, logger),
		PreferenceModel:  *preferencemodel.New(db, logger),
	}
}
