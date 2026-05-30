package createcontact

import (
	"context"

	"go-nggong/internal/dog/domain"
	dcc "go-nggong/internal/dog/domain/createcontact"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dcc.Request, presenter Presenter) error {
	saved, err := u.Gateway.Save(ctx, domain.ContactRecord{
		DogID:     request.DogID,
		Phone:     request.Phone,
		Address:   request.Address,
		Emergency: request.Emergency,
	})
	if err != nil {
		return err
	}
	presenter.Present(dcc.Result{Contact: saved})
	return nil
}
