package getall

import (
	"context"

	"go-nggong/internal/dog/domain"
)

type Gateway interface {
	FindAll(ctx context.Context) ([]domain.DogRecord, error)
}
