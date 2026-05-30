package createcontact

import "go-nggong/internal/dog/domain/createcontact"

type Presenter interface {
	Present(result createcontact.Result)
}
