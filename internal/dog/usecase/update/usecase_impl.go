package update

import (
	"context"

	"go-nggong/internal/dog/domain"
	dupdate "go-nggong/internal/dog/domain/update"
)

type UseCaseImpl struct {
	Gateway Gateway
}

func NewUseCase(g Gateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: g}
}

func (u *UseCaseImpl) Execute(ctx context.Context, request dupdate.Request, presenter Presenter) error {
	id := request.ID
	updated, err := u.Gateway.Update(ctx, domain.DogRecord{
		ID:    &id,
		Name:  request.Name,
		Email: request.Email,
		Role:  request.Role,
	})
	if err != nil {
		return err
	}
	if updated == nil {
		presenter.PresentNotFound(request.ID)
		return nil
	}
	presenter.Present(dupdate.Result{Dog: *updated})
	return nil
}
