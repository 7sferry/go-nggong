package updateprice

import (
	"context"

	"go-nggong/internal/dogprice/domain"

	"github.com/shopspring/decimal"
)

type Gateway interface {
	UpdateByDogID(ctx context.Context, dogID int, price decimal.Decimal, priceType string) (*domain.DogPriceRecord, error)
}
