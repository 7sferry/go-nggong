package create

import (
	"context"

	"go-nggong/internal/dogprice/domain/create"
)

type UseCase interface {
	Execute(ctx context.Context, request create.Request, presenter Presenter) error
}
