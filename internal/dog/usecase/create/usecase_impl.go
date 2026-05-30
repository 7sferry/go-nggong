package create

import (
	"context"

	"go-nggong/internal/dog/domain"
	dcreate "go-nggong/internal/dog/domain/create"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dcreate.Request, presenter Presenter) error {
	saved, err := u.Gateway.Save(ctx, domain.DogRecord{
		Name:  request.Name,
		Email: request.Email,
		Role:  request.Role,
	})
	if err != nil {
		return err
	}
	presenter.Present(dcreate.Result{Dog: saved})
	return nil
}
