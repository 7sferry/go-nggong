package getcontact

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	dgetcontact "go-nggong/internal/dog/domain/getcontact"
	"go-nggong/internal/dog/usecase/getcontact"
)

type Controller struct {
	UseCase getcontact.UseCase
	Log     *slog.Logger
}

func New(uc getcontact.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	dogIDStr := r.PathValue("dogId")
	dogID, err := strconv.Atoi(dogIDStr)
	if err != nil {
		http.Error(w, "invalid dogId", http.StatusBadRequest)
		return
	}
	presenter := &WebPresenter{DogID: dogID}
	if err := c.UseCase.Execute(r.Context(), dgetcontact.Request{DogID: dogID}, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "getContact failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if presenter.NotFound {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("No contact found for dog %d", dogID)})
		return
	}
	_ = json.NewEncoder(w).Encode(presenter.Response)
}
