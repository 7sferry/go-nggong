package updateprice

import (
	dupdate "go-nggong/internal/dogprice/domain/updateprice"
)

type WebPresenter struct {
	Response dupdate.Response
	NotFound bool
}

func (p *WebPresenter) Present(result dupdate.Result) {
	p.Response = dupdate.Response{
		ID:        result.Price.ID,
		DogID:     result.Price.DogID,
		Price:     result.Price.Price,
		PriceType: result.Price.PriceType,
	}
}

func (p *WebPresenter) PresentNotFound(_ int) {
	p.NotFound = true
}
