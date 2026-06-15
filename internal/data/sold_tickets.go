package data

type SoldTickets struct {
	Id          string
	UserId      string
	ArtistId    int
	ConcertDate string
	TicketType  string
	Qty         int
	Vat         *float64
	Amt         *float64
	Location    string
	BookingFee  *float64
}
