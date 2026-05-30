package updateprice

import (
	"context"

	"go-nggong/internal/dogprice/domain/updateprice"
)

type UseCase interface {
	Execute(ctx context.Context, request updateprice.Request, presenter Presenter) error
}
