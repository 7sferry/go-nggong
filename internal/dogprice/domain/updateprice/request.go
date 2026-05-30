package updateprice

import "github.com/shopspring/decimal"

type Request struct {
	DogID     int
	Price     decimal.Decimal
	PriceType string
}
