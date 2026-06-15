package soldticketsmodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

/*
id uuid NOT NULL PRIMARY KEY,
    userId uuid NOT NULL,
    artistId uuid NOT NULL,
    version integer NOT NULL DEFAULT 1,
*/

const (
	sourceGetAll = "Get all f(n) under soldticketsmodel pkg"
	sourceGet    = "Get f(n) under soldticketsmodel pkg"
)

func (s *SoldTicketsModel) Get(artistId int, userId, location, date string) (data.SoldTickets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	query := "SELECT artistId, concertDate, ticketType, qty, vat, amt, location, bookingFee FROM sold_tickets WHERE userId = $1 AND artistId = $2 AND location = $3 AND date = $4"

	row := s.db.QueryRow(ctx, query, userId, artistId)

	var sold data.SoldTickets

	if err := row.Scan(&sold.ArtistId, &sold.ConcertDate, &sold.TicketType, &sold.Qty, &sold.Vat, &sold.Amt, &sold.Location, &sold.BookingFee); err != nil {
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
			"Source": sourceGet,
		})
		return data.SoldTickets{}, e
	}
	return sold, nil
}

func (s *SoldTicketsModel) GetAll(userId string) ([]data.SoldTickets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	query := `SELECT artistId, concertDate, ticketType, qty, vat, amt, location, bookingFee FROM sold_tickets WHERE userId = $1`

	rows, err := s.db.Query(ctx, query, userId)
	if err != nil {
		var e error
		switch {
		case errors.Is(err, context.Canceled):
			e = helper.WrapError("Query execution error: context canceled", err)
		case errors.Is(err, context.DeadlineExceeded):
			e = helper.WrapError("Query execution error: deadline exceeded", err)
		case errors.Is(err, pgx.ErrTxClosed):
			e = helper.WrapError("Query execution error: transaction closed", err)
		case errors.Is(err, pgx.ErrNoRows):
			e = helper.WrapError("No favorite available for this user", err)
		default:
			e = helper.WrapError("Query execution error", err)
		}
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceGetAll,
		})
		return nil, e
	}

	var soldTickets []data.SoldTickets

	for rows.Next() {
		var soldTicket data.SoldTickets

		if err := rows.Scan(&soldTicket.ArtistId, &soldTicket.ConcertDate, &soldTicket.TicketType, &soldTicket.Qty, &soldTicket.Vat, &soldTicket.Amt, &soldTicket.Location, &soldTicket.BookingFee); err != nil {
			var e error
			switch {
			case errors.Is(err, context.Canceled):
				e = helper.WrapError("Query execution error: context canceled", err)
			case errors.Is(err, context.DeadlineExceeded):
				e = helper.WrapError("Query execution error: deadline exceeded", err)
			case errors.Is(err, pgx.ErrTxClosed):
				e = helper.WrapError("Query execution error: transaction closed", err)
			case errors.Is(err, pgx.ErrNoRows):
				e = helper.WrapError("No favorite available for this user", err)
			default:
				e = helper.WrapError("Query execution error", err)
			}
			s.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceGetAll,
			})
			return nil, e
		}

		soldTickets = append(soldTickets, soldTicket)
	}

	if err := rows.Err(); err != nil {
		var e error
		switch {
		case errors.Is(err, context.Canceled):
			e = helper.WrapError("Query execution error: context canceled", err)
		case errors.Is(err, context.DeadlineExceeded):
			e = helper.WrapError("Query execution error: deadline exceeded", err)
		case errors.Is(err, pgx.ErrTxClosed):
			e = helper.WrapError("Query execution error: transaction closed", err)
		case errors.Is(err, pgx.ErrNoRows):
			e = helper.WrapError("No favorite available for this user", err)
		default:
			e = helper.WrapError("Query execution error", err)
		}
		s.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceGetAll,
		})
		return nil, e
	}
	return soldTickets, nil
}
