package getall

import "go-nggong/internal/dog/domain/getall"

type Presenter interface {
	Present(result getall.Result)
}
