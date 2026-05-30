package update

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"go-nggong/internal/dog/controller"
	dupdate "go-nggong/internal/dog/domain/update"
	"go-nggong/internal/dog/usecase/update"
)

type Controller struct {
	UseCase update.UseCase
	Log     *slog.Logger
}

func New(uc update.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

type body struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var b body
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	c.Log.InfoContext(r.Context(), "updateDog", "id", id, "authContext", controller.ExtractAuthContext(r))
	presenter := &WebPresenter{}
	if err := c.UseCase.Execute(r.Context(), dupdate.Request{ID: id, Name: b.Name, Email: b.Email, Role: b.Role}, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "updateDog failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if presenter.NotFound {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Dog not found"})
		return
	}
	_ = json.NewEncoder(w).Encode(presenter.Response)
}
