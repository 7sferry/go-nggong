package updatecontact

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"go-nggong/internal/dog/controller"
	duc "go-nggong/internal/dog/domain/updatecontact"
	"go-nggong/internal/dog/usecase/updatecontact"
)

type Controller struct {
	UseCase updatecontact.UseCase
	Log     *slog.Logger
}

func New(uc updatecontact.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

type body struct {
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Emergency string `json:"emergency"`
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	dogIDStr := r.PathValue("dogId")
	dogID, err := strconv.Atoi(dogIDStr)
	if err != nil {
		http.Error(w, "invalid dogId", http.StatusBadRequest)
		return
	}
	var b body
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	c.Log.InfoContext(r.Context(), "updateContact", "dogId", dogID, "authContext", controller.ExtractAuthContext(r))
	presenter := &WebPresenter{}
	req := duc.Request{DogID: dogID, Phone: b.Phone, Address: b.Address, Emergency: b.Emergency}
	if err := c.UseCase.Execute(r.Context(), req, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "updateContact failed", "err", err)
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
