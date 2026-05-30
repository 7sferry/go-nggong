package getall

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"go-nggong/internal/dog/controller"
	dgetall "go-nggong/internal/dog/domain/getall"
	"go-nggong/internal/dog/usecase/getall"
)

type Controller struct {
	UseCase getall.UseCase
	Log     *slog.Logger
}

func New(uc getall.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	c.Log.InfoContext(r.Context(), "getAllDogs", "authContext", controller.ExtractAuthContext(r))
	presenter := &WebPresenter{}
	if err := c.UseCase.Execute(r.Context(), dgetall.Request{}, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "getAllDogs failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(presenter.Response)
}
