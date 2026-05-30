package getall

import (
	"context"

	dgetall "go-nggong/internal/dogprice/domain/getall"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, _ dgetall.Request, presenter Presenter) error {
	prices, err := u.Gateway.FindAll(ctx)
	if err != nil {
		return err
	}
	presenter.Present(dgetall.Result{Prices: prices})
	return nil
}
