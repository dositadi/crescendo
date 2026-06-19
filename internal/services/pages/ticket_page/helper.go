package ticketpage

import (
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

func (t *TicketPage) atoi(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		t.logger.PrintError("Atoi conversion error: Not a valid number", map[string]string{
			"Source": sourceT,
		})
		panic("Not a valid number")
	}
	return out
}

func (t *TicketPage) detailPageFuncMap() template.FuncMap {
	return template.FuncMap{
		"GetDay": func(date string) string {
			out := strings.Split(date, "-")
			return out[0]
		},
		"GetMonth": func(date string) string {
			out := strings.Split(date, "-")
			num := t.atoi(out[1])

			switch num {
			case 1:
				return "JAN"
			case 2:
				return "FEB"
			case 3:
				return "MAR"
			case 4:
				return "APR"
			case 5:
				return "MAY"
			case 6:
				return "JUN"
			case 7:
				return "JUL"
			case 8:
				return "AUG"
			case 9:
				return "SEPT"
			case 10:
				return "OCT"
			case 11:
				return "NOV"
			case 12:
				return "DEC"
			default:
				return "JAN"
			}
		},
		"GetYear": func(date string) string {
			out := strings.Split(date, "-")
			return out[2]
		},
		"CleanCityName": func(city string) string {
			city = strings.ReplaceAll(city, "-", " ")
			city = strings.ReplaceAll(city, "_", "-")
			citySlice := strings.Split(city, " ")

			for i, c := range citySlice {
				v := []rune(c)
				citySlice[i] = string(unicode.ToUpper(v[0])) + string(v[1:])
			}

			return strings.Join(citySlice, ", ")
		},
		"RandomValues": func() int {
			return rand.Intn(500)
		},
		"GetArtistImage": func(artistId int, info string) string {
			artistInfo := t.client.Get()[artistId]

			switch info {
			case "image":
				return artistInfo.Image
			case "name":
				return artistInfo.Name
			}
			return ""
		},
		"DeReferenceValue": func(val *float64) float64 {
			if val != nil {
				return *val
			}
			return 0.0
		},
		"FormatTime": func(time time.Time) string {
			return time.Format("Jan 2, 2006")
		},
	}
}

func (t *TicketPage) getUser() data.User {
	val := t.request.Context().Value(utils.USER_ID_KEY)

	if user, ok := val.(data.User); ok {
		return user
	}
	return data.User{}
}
