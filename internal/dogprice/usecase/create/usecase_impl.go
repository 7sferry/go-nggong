package create

import (
	"context"

	"go-nggong/internal/dogprice/domain"
	dcreate "go-nggong/internal/dogprice/domain/create"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dcreate.Request, presenter Presenter) error {
	saved, err := u.Gateway.Save(ctx, domain.DogPriceRecord{
		DogID:     request.DogID,
		Price:     request.Price,
		PriceType: request.PriceType,
	})
	if err != nil {
		return err
	}
	presenter.Present(dcreate.Result{Price: saved})
	return nil
}
