package getcontact

import (
	"context"

	"go-nggong/internal/dog/domain/getcontact"
)

type UseCase interface {
	Execute(ctx context.Context, request getcontact.Request, presenter Presenter) error
}
