package getall

import (
	"context"

	"go-nggong/internal/dog/domain"
	dogetall "go-nggong/internal/dog/domain/getall"
	"go-nggong/internal/dog/usecase/gateway"
)

type UseCaseImpl struct {
	Gateway         Gateway
	DogPriceGateway gateway.DogPriceGateway
}

func NewUseCase(g Gateway, dpg gateway.DogPriceGateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g, DogPriceGateway: dpg}
}

func (u *UseCaseImpl) Execute(ctx context.Context, _ dogetall.Request, presenter Presenter) error {
	dogs, err := u.Gateway.FindAll(ctx)
	if err != nil {
		return err
	}
	enriched := make([]domain.DogWithPrice, 0, len(dogs))
	for _, d := range dogs {
		entry := domain.DogWithPrice{Dog: d}
		if d.ID != nil {
			if pi, perr := u.DogPriceGateway.FindPriceByDogID(ctx, *d.ID); perr == nil && pi != nil {
				price := pi.Price
				ptype := pi.PriceType
				entry.Price = &price
				entry.PriceType = &ptype
			}
		}
		enriched = append(enriched, entry)
	}
	presenter.Present(dogetall.Result{Dogs: enriched})
	return nil
}
