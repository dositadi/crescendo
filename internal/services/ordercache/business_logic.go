package ordercache

import "math"

func GetTicketPrice(ticketType string) float64 {
	switch ticketType {
	case string(GENERAL):
		return Round(float64(GENERAL_AMT))
	case string(RESERVED):
		return Round(float64(RESERVED_AMT))
	case string(VIP):
		return Round(float64(VIP_AMT))
	default:
		return 0
	}
}

func Round[T ~float32 | ~float64](val T) T {
	return T(math.Round(float64(val*100)) / 100)
}

func TotalTicketAmount(ticketPrice float64, quantity int) float64 {
	amt := ticketPrice * float64(quantity)
	return Round(amt)
}

func TotalBookingFee(fee float64, qty int) float64 {
	amt := fee * float64(qty)
	return Round(amt)
}

func VatAmount(vatRate, totalTicketAmount, totalBookingFee float64) float64 {
	totalPrice := totalBookingFee + totalTicketAmount
	vatAmt := totalPrice * (vatRate / (100 + vatRate))
	return Round(vatAmt)
}

func GrandTotal(totalTicketAmount, totalBookingFee, vatAmt float64) float64 {
	total := totalBookingFee + totalTicketAmount + vatAmt
	return Round(total)
}
