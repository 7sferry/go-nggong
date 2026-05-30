package createcontact

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"go-nggong/internal/dog/controller"
	dcc "go-nggong/internal/dog/domain/createcontact"
	"go-nggong/internal/dog/usecase/createcontact"
)

type Controller struct {
	UseCase createcontact.UseCase
	Log     *slog.Logger
}

func New(uc createcontact.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	var req dcc.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	c.Log.InfoContext(r.Context(), "createContact", "dogId", req.DogID, "authContext", controller.ExtractAuthContext(r))
	presenter := &WebPresenter{}
	if err := c.UseCase.Execute(r.Context(), req, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "createContact failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(presenter.Response)
}
