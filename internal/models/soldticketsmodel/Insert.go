package soldticketsmodel

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceInsert = "Insert f(n) under soldticketsmodel pkg"
)

var CONFLICT_ERR = fmt.Errorf("Sold Ticket exists already exists already")

func (s *SoldTicketsModel) Insert(soldTicket data.SoldTickets) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	defer tx.Rollback(ctx)
	if err != nil {
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
			"Source": sourceInsert,
		})
		return e
	}

	exists, err := s.Exists(soldTicket.UserId, soldTicket.ConcertDate, soldTicket.Location, soldTicket.ArtistId)
	if err != nil {
		return err
	}
	if exists {
		s.logger.PrintError(CONFLICT_ERR.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return CONFLICT_ERR
	}

	query := "INSERT INTO favorites (id, userId, artistId, userContactFName, userContactLName, userContactEmail, concertDate, ticketType, qty, vat, amt, location, bookingFee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"

	_, err = tx.Exec(ctx, query, soldTicket.Id, soldTicket.UserId, soldTicket.ArtistId, soldTicket.UserContactFName, soldTicket.UserContactLName, soldTicket.UserContactEmail, soldTicket.ConcertDate, soldTicket.TicketType, soldTicket.Qty, soldTicket.Vat, soldTicket.Amt, soldTicket.Location, soldTicket.BookingFee)
	if err != nil {
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
			"Source": sourceInsert,
		})
		return e
	}

	if err = tx.Commit(ctx); err != nil {
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
			"Source": sourceInsert,
		})
		return e
	}

	return nil
}
