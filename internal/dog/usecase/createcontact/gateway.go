package createcontact

import (
	"context"

	"go-nggong/internal/dog/domain"
)

type Gateway interface {
	Save(ctx context.Context, contact domain.ContactRecord) (domain.ContactRecord, error)
}
