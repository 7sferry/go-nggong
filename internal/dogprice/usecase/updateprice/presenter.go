package updateprice

import "go-nggong/internal/dogprice/domain/updateprice"

type Presenter interface {
	Present(result updateprice.Result)
	PresentNotFound(dogID int)
}
