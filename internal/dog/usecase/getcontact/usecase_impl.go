package getcontact

import (
	"context"

	dgetcontact "go-nggong/internal/dog/domain/getcontact"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dgetcontact.Request, presenter Presenter) error {
	contact, err := u.Gateway.FindByDogID(ctx, request.DogID)
	if err != nil {
		return err
	}
	if contact == nil {
		presenter.PresentNotFound(request.DogID)
		return nil
	}
	presenter.Present(dgetcontact.Result{Contact: *contact})
	return nil
}
