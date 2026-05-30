package updatecontact

import (
	duc "go-nggong/internal/dog/domain/updatecontact"
)

type WebPresenter struct {
	Response duc.Response
	NotFound bool
}

func (p *WebPresenter) Present(result duc.Result) {
	c := result.Contact
	p.Response = duc.Response{
		ID:        c.ID,
		DogID:     c.DogID,
		Phone:     c.Phone,
		Address:   c.Address,
		Emergency: c.Emergency,
	}
}

func (p *WebPresenter) PresentNotFound(_ int) {
	p.NotFound = true
}
