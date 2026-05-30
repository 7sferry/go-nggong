package getbyid

import (
	"context"

	"go-nggong/internal/dog/domain"
	dogbyid "go-nggong/internal/dog/domain/getbyid"
	"go-nggong/internal/dog/usecase/gateway"
)

type UseCaseImpl struct {
	Gateway         Gateway
	DogPriceGateway gateway.DogPriceGateway
}

func NewUseCase(g Gateway, dpg gateway.DogPriceGateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g, DogPriceGateway: dpg}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dogbyid.Request, presenter Presenter) error {
	dog, err := u.Gateway.FindByID(ctx, request.ID)
	if err != nil {
		return err
	}
	if dog == nil {
		presenter.PresentNotFound(request.ID)
		return nil
	}
	entry := domain.DogWithPrice{Dog: *dog}
	if dog.ID != nil {
		if pi, perr := u.DogPriceGateway.FindPriceByDogID(ctx, *dog.ID); perr == nil && pi != nil {
			price := pi.Price
			ptype := pi.PriceType
			entry.Price = &price
			entry.PriceType = &ptype
		}
	}
	presenter.Present(dogbyid.Result{DogWithPrice: entry})
	return nil
}
