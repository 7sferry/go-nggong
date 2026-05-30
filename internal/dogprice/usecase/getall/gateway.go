package getall

import (
	"context"

	"go-nggong/internal/dogprice/domain"
)

type Gateway interface {
	FindAll(ctx context.Context) ([]domain.DogPriceRecord, error)
}
