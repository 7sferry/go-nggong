package domain

import "github.com/shopspring/decimal"

type DogPriceRecord struct {
	ID        *int
	DogID     int
	Price     decimal.Decimal
	PriceType string
}
