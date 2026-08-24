package soldticketsmodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SoldTicketsModel struct {
	db     *pgxpool.Pool
	logger jsonlog.Logger
}

func New(db *pgxpool.Pool, logger jsonlog.Logger) *SoldTicketsModel {
	return &SoldTicketsModel{
		logger: logger,
		db:     db,
	}
}
