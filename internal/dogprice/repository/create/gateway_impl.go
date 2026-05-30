package create

import (
	"context"
	"database/sql"

	"go-nggong/internal/dogprice/domain"
)

type GatewayImpl struct {
	DB *sql.DB
}

func New(db *sql.DB) *GatewayImpl {
	return &GatewayImpl{DB: db}
}

func (g *GatewayImpl) Save(ctx context.Context, price domain.DogPriceRecord) (domain.DogPriceRecord, error) {
	var id int
	err := g.DB.QueryRowContext(ctx,
		"INSERT INTO dog_prices (dog_id, price, price_type) VALUES ($1, $2, $3) RETURNING id",
		price.DogID, price.Price, price.PriceType).Scan(&id)
	if err != nil {
		return domain.DogPriceRecord{}, err
	}
	price.ID = &id
	return price, nil
}
