package updatecontact

import (
	"context"

	"go-nggong/internal/dog/domain"
	duc "go-nggong/internal/dog/domain/updatecontact"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request duc.Request, presenter Presenter) error {
	updated, err := u.Gateway.UpdateByDogID(ctx, domain.ContactRecord{
		DogID:     request.DogID,
		Phone:     request.Phone,
		Address:   request.Address,
		Emergency: request.Emergency,
	})
	if err != nil {
		return err
	}
	if updated == nil {
		presenter.PresentNotFound(request.DogID)
		return nil
	}
	presenter.Present(duc.Result{Contact: *updated})
	return nil
}
