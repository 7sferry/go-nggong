package create

import "go-nggong/internal/dogprice/domain/create"

type Presenter interface {
	Present(result create.Result)
}
