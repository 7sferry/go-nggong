package getall

import "go-nggong/internal/dogprice/domain/getall"

type Presenter interface {
	Present(result getall.Result)
}
