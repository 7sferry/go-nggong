package updateprice

import (
	"context"

	dupdate "go-nggong/internal/dogprice/domain/updateprice"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dupdate.Request, presenter Presenter) error {
	updated, err := u.Gateway.UpdateByDogID(ctx, request.DogID, request.Price, request.PriceType)
	if err != nil {
		return err
	}
	if updated == nil {
		presenter.PresentNotFound(request.DogID)
		return nil
	}
	presenter.Present(dupdate.Result{Price: *updated})
	return nil
}
