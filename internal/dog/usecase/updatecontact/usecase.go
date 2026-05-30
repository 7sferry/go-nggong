package updatecontact

import (
	"context"

	"go-nggong/internal/dog/domain/updatecontact"
)

type UseCase interface {
	Execute(ctx context.Context, request updatecontact.Request, presenter Presenter) error
}
