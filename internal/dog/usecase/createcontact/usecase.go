package createcontact

import (
	"context"

	"go-nggong/internal/dog/domain/createcontact"
)

type UseCase interface {
	Execute(ctx context.Context, request createcontact.Request, presenter Presenter) error
}
