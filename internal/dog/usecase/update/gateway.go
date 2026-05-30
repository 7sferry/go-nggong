package update

import (
	"context"

	"go-nggong/internal/dog/domain"
)

type Gateway interface {
	Update(ctx context.Context, dog domain.DogRecord) (*domain.DogRecord, error)
}
