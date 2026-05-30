package updateprice

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	dupdate "go-nggong/internal/dogprice/domain/updateprice"
	"go-nggong/internal/dogprice/usecase/updateprice"

	"github.com/shopspring/decimal"
)

type Controller struct {
	UseCase updateprice.UseCase
	Log     *slog.Logger
}

func New(uc updateprice.UseCase, log *slog.Logger) *Controller {
	return &Controller{UseCase: uc, Log: log}
}

type body struct {
	Price     decimal.Decimal `json:"price"`
	PriceType string          `json:"priceType"`
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
	presenter := &WebPresenter{}
	req := dupdate.Request{DogID: dogID, Price: b.Price, PriceType: b.PriceType}
	if err := c.UseCase.Execute(r.Context(), req, presenter); err != nil {
		c.Log.ErrorContext(r.Context(), "updatePrice failed", "err", err)
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
