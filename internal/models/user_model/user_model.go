package usermodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserModel struct {
	db     *pgxpool.Pool
	logger jsonlog.Logger
}

const (
	timeOut = 5
)

func New(db *pgxpool.Pool, logger jsonlog.Logger) *UserModel {
	return &UserModel{
		db:     db,
		logger: logger,
	}
}
