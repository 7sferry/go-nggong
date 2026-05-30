package getprice

import (
	"context"

	"go-nggong/internal/dogprice/domain/getprice"
)

type UseCase interface {
	Execute(ctx context.Context, request getprice.Request, presenter Presenter) error
}
