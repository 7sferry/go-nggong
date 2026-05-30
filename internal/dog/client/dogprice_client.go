package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go-nggong/internal/dog/usecase/gateway"

	"github.com/shopspring/decimal"
)

type DogPriceClient struct {
	BaseURL string
	HTTP    *http.Client
}

func New(baseURL string) *DogPriceClient {
	return &DogPriceClient{
		BaseURL: baseURL,
		HTTP:    &http.Client{Timeout: 5 * time.Second},
	}
}

type priceResponse struct {
	DogID     int             `json:"dogId"`
	Price     decimal.Decimal `json:"price"`
	PriceType string          `json:"priceType"`
}

func (c *DogPriceClient) FindPriceByDogID(ctx context.Context, dogID int) (*gateway.DogPriceInfo, error) {
	url := fmt.Sprintf("%s/prices/%d", c.BaseURL, dogID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}
	var pr priceResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return nil, err
	}
	return &gateway.DogPriceInfo{Price: pr.Price, PriceType: pr.PriceType}, nil
}
