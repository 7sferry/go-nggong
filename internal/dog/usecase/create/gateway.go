package create

import (
	"context"

	"go-nggong/internal/dog/domain"
)

type Gateway interface {
	Save(ctx context.Context, dog domain.DogRecord) (domain.DogRecord, error)
}
