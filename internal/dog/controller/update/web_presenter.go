package update

import (
	dupdate "go-nggong/internal/dog/domain/update"
)

type WebPresenter struct {
	Response dupdate.Response
	NotFound bool
}

func (p *WebPresenter) Present(result dupdate.Result) {
	p.Response = dupdate.Response{
		ID:    result.Dog.ID,
		Name:  result.Dog.Name,
		Email: result.Dog.Email,
		Role:  result.Dog.Role,
	}
}

func (p *WebPresenter) PresentNotFound(_ int) {
	p.NotFound = true
}
