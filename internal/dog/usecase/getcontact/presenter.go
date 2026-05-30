package getcontact

import "go-nggong/internal/dog/domain/getcontact"

type Presenter interface {
	Present(result getcontact.Result)
	PresentNotFound(dogID int)
}
