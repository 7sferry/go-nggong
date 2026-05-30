package getbyid

import "go-nggong/internal/dog/domain/getbyid"

type Presenter interface {
	Present(result getbyid.Result)
	PresentNotFound(id int)
}
