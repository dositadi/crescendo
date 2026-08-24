package favoritemodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FavoriteModel struct {
	db     *pgxpool.Pool
	logger jsonlog.Logger
}

func New(db *pgxpool.Pool, logger jsonlog.Logger) *FavoriteModel {
	return &FavoriteModel{
		logger: logger,
		db:     db,
	}
}
