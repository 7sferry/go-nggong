package getprice

import (
	"context"

	dgetprice "go-nggong/internal/dogprice/domain/getprice"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dgetprice.Request, presenter Presenter) error {
	price, err := u.Gateway.FindByDogID(ctx, request.DogID)
	if err != nil {
		return err
	}
	if price == nil {
		presenter.PresentNotFound(request.DogID)
		return nil
	}
	presenter.Present(dgetprice.Result{Price: *price})
	return nil
}
