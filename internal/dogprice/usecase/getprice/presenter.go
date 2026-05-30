package getprice

import "go-nggong/internal/dogprice/domain/getprice"

type Presenter interface {
	Present(result getprice.Result)
	PresentNotFound(dogID int)
}
