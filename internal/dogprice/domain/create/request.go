package create

import "github.com/shopspring/decimal"

type Request struct {
	DogID     int             `json:"dogId"`
	Price     decimal.Decimal `json:"price"`
	PriceType string          `json:"priceType"`
}
