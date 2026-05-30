package updateprice

import "github.com/shopspring/decimal"

type Response struct {
	ID        *int            `json:"id"`
	DogID     int             `json:"dogId"`
	Price     decimal.Decimal `json:"price"`
	PriceType string          `json:"priceType"`
}
