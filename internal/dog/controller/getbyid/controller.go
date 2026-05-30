package getbyid

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"go-nggong/internal/dog/controller"
	dgetbyid "go-nggong/internal/dog/domain/getbyid"
	"go-nggong/internal/dog/usecase/getbyid"
)

type Controller struct {
	UseCase getbyid.UseCase
	Log     *slog.Logger
}

func New(uc getbyid.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	c.Log.InfoContext(r.Context(), "getDogByID", "id", id, "authContext", controller.ExtractAuthContext(r))
	presenter := &WebPresenter{}
	if err := c.UseCase.Execute(r.Context(), dgetbyid.Request{ID: id}, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "getDogByID failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	writeResponse(w, presenter)
}

func writeResponse(w http.ResponseWriter, p *WebPresenter) {
	w.Header().Set("Content-Type", "application/json")
	if p.NotFound {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Dog not found"})
		return
	}
	_ = json.NewEncoder(w).Encode(p.Response)
}
