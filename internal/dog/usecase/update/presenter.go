package update

import "go-nggong/internal/dog/domain/update"

type Presenter interface {
	Present(result update.Result)
	PresentNotFound(id int)
}
