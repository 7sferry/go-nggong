package getall

import "github.com/shopspring/decimal"

type Response struct {
	Prices []PriceEntry `json:"prices"`
}

type PriceEntry struct {
	DogID     int             `json:"dogId"`
	Price     decimal.Decimal `json:"price"`
	PriceType string          `json:"priceType"`
}
