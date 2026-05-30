package getall

import (
	"encoding/json"
	"log/slog"
	"net/http"

	dgetall "go-nggong/internal/dogprice/domain/getall"
	"go-nggong/internal/dogprice/usecase/getall"
)

type Controller struct {
	UseCase getall.UseCase
	Log     *slog.Logger
}

func New(uc getall.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	presenter := &WebPresenter{}
	if err := c.UseCase.Execute(r.Context(), dgetall.Request{}, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "getAllPrices failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(presenter.Response)
}
