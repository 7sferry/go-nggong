package updatecontact

import "go-nggong/internal/dog/domain/updatecontact"

type Presenter interface {
	Present(result updatecontact.Result)
	PresentNotFound(dogID int)
}
