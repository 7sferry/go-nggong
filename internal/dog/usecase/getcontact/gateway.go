package getcontact

import (
	"context"

	"go-nggong/internal/dog/domain"
)

type Gateway interface {
	FindByDogID(ctx context.Context, dogID int) (*domain.ContactRecord, error)
}
