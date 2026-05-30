package create

import (
	"context"

	"go-nggong/internal/dog/domain/create"
)

type UseCase interface {
	Execute(ctx context.Context, request create.Request, presenter Presenter) error
}
