package getbyid

import (
	dgetbyid "go-nggong/internal/dog/domain/getbyid"
)

type WebPresenter struct {
	Response dgetbyid.Response
	NotFound bool
}

func (p *WebPresenter) Present(result dgetbyid.Result) {
	entry := result.DogWithPrice
	var price any = "unavailable"
	var ptype any = "unavailable"
	if entry.Price != nil {
		price = *entry.Price
	}
	if entry.PriceType != nil {
		ptype = *entry.PriceType
	}
	p.Response = dgetbyid.Response{
		Dog: dgetbyid.DogWithPrice{
			ID:        entry.Dog.ID,
			Name:      entry.Dog.Name,
			Email:     entry.Dog.Email,
			Role:      entry.Dog.Role,
			Price:     price,
			PriceType: ptype,
		},
	}
}

func (p *WebPresenter) PresentNotFound(_ int) {
	p.NotFound = true
}
