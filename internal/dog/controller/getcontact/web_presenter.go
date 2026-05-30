package getcontact

import (
	dgetcontact "go-nggong/internal/dog/domain/getcontact"
)

type WebPresenter struct {
	DogID    int
	Response dgetcontact.Response
	NotFound bool
}

func (p *WebPresenter) Present(result dgetcontact.Result) {
	c := result.Contact
	p.Response = dgetcontact.Response{
		DogID: c.DogID,
		Contact: map[string]string{
			"phone":     c.Phone,
			"address":   c.Address,
			"emergency": c.Emergency,
		},
	}
}

func (p *WebPresenter) PresentNotFound(_ int) {
	p.NotFound = true
}
