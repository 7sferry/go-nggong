package createcontact

import (
	dcc "go-nggong/internal/dog/domain/createcontact"
)

type WebPresenter struct {
	Response dcc.Response
}

func (p *WebPresenter) Present(result dcc.Result) {
	c := result.Contact
	p.Response = dcc.Response{
		ID:        c.ID,
		DogID:     c.DogID,
		Phone:     c.Phone,
		Address:   c.Address,
		Emergency: c.Emergency,
	}
}
