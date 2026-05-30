package create

import "go-nggong/internal/dog/domain/create"

type Presenter interface {
	Present(result create.Result)
}
