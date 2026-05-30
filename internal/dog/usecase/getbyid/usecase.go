package getbyid

import (
	"context"

	"go-nggong/internal/dog/domain/getbyid"
)

type UseCase interface {
	Execute(ctx context.Context, request getbyid.Request, presenter Presenter) error
}
