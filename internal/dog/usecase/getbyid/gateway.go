package getbyid

import (
	"context"

	"go-nggong/internal/dog/domain"
)

type Gateway interface {
	FindByID(ctx context.Context, id int) (*domain.DogRecord, error)
}
