package preferencemodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PreferenceModel struct {
	db     *pgxpool.Pool
	logger jsonlog.Logger
}

func New(db *pgxpool.Pool, logger jsonlog.Logger) *PreferenceModel {
	return &PreferenceModel{
		logger: logger,
		db:     db,
	}
}
