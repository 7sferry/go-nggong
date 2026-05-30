package getprice

import "github.com/shopspring/decimal"

type Response struct {
	DogID     int             `json:"dogId"`
	Price     decimal.Decimal `json:"price"`
	PriceType string          `json:"priceType"`
}
