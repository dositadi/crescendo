package ticketpage

import (
	"fmt"
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceTI = "ticketpage.RenderTicketPage()"
)

func (t *TicketPage) RenderTicketReciept() error {
	fs := []string{
		"internal/web/static/pages/reciepts_page.html",
	}
	artistId := t.request.FormValue(utils.ARTIST_ID_KEY)
	location := t.request.FormValue(utils.LOCATION_KEY)
	date := t.request.FormValue(utils.DATE_KEY)
	path := t.request.FormValue(utils.PATH_KEY)
	fmt.Println(path)

	user := t.getUser()

	userTicket, err := t.soldTicketsModel.Get(t.atoi(artistId), user.Id, location, date)
	if err != nil {
		e := helper.WrapError("Get ticket error", err)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceTI,
		})
		return e
	}

	data := struct {
		PreviousPageUrl                                    string
		BoughtTicket                                       data.SoldTickets
		TotalTicketAmount, TotalBookingFee, TotalVatAmount float64
	}{
		PreviousPageUrl:   path,
		BoughtTicket:      userTicket,
		TotalTicketAmount: ordercache.TotalTicketAmount(ordercache.GetTicketPrice(userTicket.TicketType), userTicket.Qty),
		TotalBookingFee:   ordercache.TotalBookingFee(*userTicket.BookingFee, userTicket.Qty),
		TotalVatAmount:    ordercache.VatAmount(*userTicket.Vat, ordercache.TotalTicketAmount(ordercache.GetTicketPrice(userTicket.TicketType), userTicket.Qty), ordercache.TotalBookingFee(*userTicket.BookingFee, userTicket.Qty)),
	}

	temp, err := template.New("reciepts_page.html").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template parse error", err)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceTI,
		})
		return e
	}

	if err := temp.Execute(t.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceTI,
		})
		return e
	}
	return nil
}
