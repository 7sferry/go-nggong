package create

import (
	"encoding/json"
	"log/slog"
	"net/http"

	dcreate "go-nggong/internal/dogprice/domain/create"
	"go-nggong/internal/dogprice/usecase/create"
)

type Controller struct {
	UseCase create.UseCase
	Log     *slog.Logger
}

func New(uc create.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	var req dcreate.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	presenter := &WebPresenter{}
	if err := c.UseCase.Execute(r.Context(), req, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "createPrice failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(presenter.Response)
}
