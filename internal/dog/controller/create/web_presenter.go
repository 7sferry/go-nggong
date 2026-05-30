package create

import (
	dcreate "go-nggong/internal/dog/domain/create"
)

type WebPresenter struct {
	Response dcreate.Response
}

func (p *WebPresenter) Present(result dcreate.Result) {
	p.Response = dcreate.Response{
		ID:    result.Dog.ID,
		Name:  result.Dog.Name,
		Email: result.Dog.Email,
		Role:  result.Dog.Role,
	}
}
