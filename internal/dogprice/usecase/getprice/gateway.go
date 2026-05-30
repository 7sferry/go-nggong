package getprice

import (
	"context"

	"go-nggong/internal/dogprice/domain"
)

type Gateway interface {
	FindByDogID(ctx context.Context, dogID int) (*domain.DogPriceRecord, error)
}
