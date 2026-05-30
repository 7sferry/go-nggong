package getall

import (
	dgetall "go-nggong/internal/dog/domain/getall"
)

type WebPresenter struct {
	Response dgetall.Response
}

func (p *WebPresenter) Present(result dgetall.Result) {
	out := make([]dgetall.DogWithPrice, 0, len(result.Dogs))
	for _, d := range result.Dogs {
		var price any = "unavailable"
		var ptype any = "unavailable"
		if d.Price != nil {
			price = *d.Price
		}
		if d.PriceType != nil {
			ptype = *d.PriceType
		}
		out = append(out, dgetall.DogWithPrice{
			ID:        d.Dog.ID,
			Name:      d.Dog.Name,
			Email:     d.Dog.Email,
			Role:      d.Dog.Role,
			Price:     price,
			PriceType: ptype,
		})
	}
	p.Response = dgetall.Response{Dogs: out}
}
