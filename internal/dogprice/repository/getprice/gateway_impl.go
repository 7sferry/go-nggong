package getprice

import (
	"context"
	"database/sql"
	"errors"

	"go-nggong/internal/dogprice/domain"
)

type GatewayImpl struct {
	DB *sql.DB
}

func New(db *sql.DB) *GatewayImpl {
	return &GatewayImpl{DB: db}
}

func (g *GatewayImpl) FindByDogID(ctx context.Context, dogID int) (*domain.DogPriceRecord, error) {
	var id int
	var p domain.DogPriceRecord
	err := g.DB.QueryRowContext(ctx,
		"SELECT id, dog_id, price, price_type FROM dog_prices WHERE dog_id = $1", dogID).
		Scan(&id, &p.DogID, &p.Price, &p.PriceType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	p.ID = &id
	return &p, nil
}
