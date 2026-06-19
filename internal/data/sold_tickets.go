package data

import "time"

type SoldTickets struct {
	Id               string
	UserId           string
	UserContactFName string
	UserContactLName string
	UserContactEmail string
	ArtistId         int
	ConcertDate      string
	TicketType       string
	Qty              int
	Vat              *float64
	Amt              *float64
	Location         string
	BookingFee       *float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
