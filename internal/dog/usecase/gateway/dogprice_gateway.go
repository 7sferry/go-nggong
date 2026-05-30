package gateway

import (
	"context"

	"github.com/shopspring/decimal"
)

type DogPriceInfo struct {
	Price     decimal.Decimal
	PriceType string
}

type DogPriceGateway interface {
	FindPriceByDogID(ctx context.Context, dogID int) (*DogPriceInfo, error)
}
