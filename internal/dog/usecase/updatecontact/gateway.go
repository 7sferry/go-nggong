package updatecontact

import (
	"context"

	"go-nggong/internal/dog/domain"
)

type Gateway interface {
	UpdateByDogID(ctx context.Context, contact domain.ContactRecord) (*domain.ContactRecord, error)
}
