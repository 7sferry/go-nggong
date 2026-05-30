package getprice

import (
	dgetprice "go-nggong/internal/dogprice/domain/getprice"
)

type WebPresenter struct {
	Response dgetprice.Response
	NotFound bool
}

func (p *WebPresenter) Present(result dgetprice.Result) {
	p.Response = dgetprice.Response{
		DogID:     result.Price.DogID,
		Price:     result.Price.Price,
		PriceType: result.Price.PriceType,
	}
}

func (p *WebPresenter) PresentNotFound(_ int) {
	p.NotFound = true
}
