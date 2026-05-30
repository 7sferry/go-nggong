package updateprice

import (
	"context"
	"database/sql"
	"errors"

	"go-nggong/internal/dogprice/domain"

	"github.com/shopspring/decimal"
)

type GatewayImpl struct {
	DB *sql.DB
}

func New(db *sql.DB) *GatewayImpl {
	return &GatewayImpl{DB: db}
}

func (g *GatewayImpl) UpdateByDogID(ctx context.Context, dogID int, price decimal.Decimal, priceType string) (*domain.DogPriceRecord, error) {
	var id int
	var dp domain.DogPriceRecord
	err := g.DB.QueryRowContext(ctx,
		"UPDATE dog_prices SET price = $1, price_type = $2 WHERE dog_id = $3 RETURNING id, dog_id, price, price_type",
		price, priceType, dogID).
		Scan(&id, &dp.DogID, &dp.Price, &dp.PriceType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	dp.ID = &id
	return &dp, nil
}
