package create

import (
	dcreate "go-nggong/internal/dogprice/domain/create"
)

type WebPresenter struct {
	Response dcreate.Response
}

func (p *WebPresenter) Present(result dcreate.Result) {
	p.Response = dcreate.Response{
		ID:        result.Price.ID,
		DogID:     result.Price.DogID,
		Price:     result.Price.Price,
		PriceType: result.Price.PriceType,
	}
}
