package create

import (
	"context"

	"go-nggong/internal/dogprice/domain"
)

type Gateway interface {
	Save(ctx context.Context, price domain.DogPriceRecord) (domain.DogPriceRecord, error)
}
