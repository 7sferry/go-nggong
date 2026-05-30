package getprice

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	dgetprice "go-nggong/internal/dogprice/domain/getprice"
	"go-nggong/internal/dogprice/usecase/getprice"
)

type Controller struct {
	UseCase getprice.UseCase
	Log     *slog.Logger
}

func New(uc getprice.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	dogIDStr := r.PathValue("dogId")
	dogID, err := strconv.Atoi(dogIDStr)
	if err != nil {
		http.Error(w, "invalid dogId", http.StatusBadRequest)
		return
	}
	presenter := &WebPresenter{}
	if err := c.UseCase.Execute(r.Context(), dgetprice.Request{DogID: dogID}, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "getPrice failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if presenter.NotFound {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("No price found for dog %d", dogID)})
		return
	}
	_ = json.NewEncoder(w).Encode(presenter.Response)
}
