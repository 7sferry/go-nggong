package getall

import (
	dgetall "go-nggong/internal/dogprice/domain/getall"
)

type WebPresenter struct {
	Response dgetall.Response
}

func (p *WebPresenter) Present(result dgetall.Result) {
	entries := make([]dgetall.PriceEntry, 0, len(result.Prices))
	for _, pr := range result.Prices {
		entries = append(entries, dgetall.PriceEntry{
			DogID:     pr.DogID,
			Price:     pr.Price,
			PriceType: pr.PriceType,
		})
	}
	p.Response = dgetall.Response{Prices: entries}
}
