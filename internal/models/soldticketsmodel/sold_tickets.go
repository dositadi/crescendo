package soldticketsmodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5"
)

type SoldTicketsModel struct {
	db     *pgx.Conn
	logger jsonlog.Logger
}

func New(db *pgx.Conn, logger jsonlog.Logger) *SoldTicketsModel {
	return &SoldTicketsModel{
		logger: logger,
		db:     db,
	}
}
