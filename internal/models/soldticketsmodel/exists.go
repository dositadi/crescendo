package soldticketsmodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceExists = "Exists f(n) under soldticketsmodel"
	timeout      = 5
)

func (s *SoldTicketsModel) Exists(userId, date, location string, artistId int) (bool, error) {
	//  Remember to create an index for the artist id
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	query := `SELECT EXISTS (SELECT 1 FROM preferences WHERE userId = $1 AND concertDate = $2 AND location = $3 AND artistId = $4)`

	row := s.db.QueryRow(ctx, query, userId, date, location, artistId)

	var exists bool

	if err := row.Scan(&exists); err != nil {
		var e error
		switch {
		case errors.Is(err, context.Canceled):
			e = helper.WrapError("Query execution error: context canceled", err)
		case errors.Is(err, context.DeadlineExceeded):
			e = helper.WrapError("Query execution error: deadline exceeded", err)
		case errors.Is(err, pgx.ErrTxClosed):
			e = helper.WrapError("Query execution error: transaction closed", err)
		default:
			e = helper.WrapError("Query execution error", err)
		}
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceExists,
		})
		return false, e
	}
	return exists, nil
}
