package update

import (
	"context"

	"go-nggong/internal/dog/domain/update"
)

type UseCase interface {
	Execute(ctx context.Context, request update.Request, presenter Presenter) error
}
