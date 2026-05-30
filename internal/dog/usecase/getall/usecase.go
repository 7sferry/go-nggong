package getall

import (
	"context"

	"go-nggong/internal/dog/domain/getall"
)

type UseCase interface {
	Execute(ctx context.Context, request getall.Request, presenter Presenter) error
}
