package domain

import "github.com/shopspring/decimal"

type DogWithPrice struct {
	Dog       DogRecord
	Price     *decimal.Decimal
	PriceType *string
}
