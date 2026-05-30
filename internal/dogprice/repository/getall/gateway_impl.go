package getall

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

func (g *GatewayImpl) FindAll(ctx context.Context) ([]domain.DogPriceRecord, error) {
	rows, err := g.DB.QueryContext(ctx, "SELECT id, dog_id, price, price_type FROM dog_prices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var prices []domain.DogPriceRecord
	for rows.Next() {
		var id int
		var p domain.DogPriceRecord
		if err := rows.Scan(&id, &p.DogID, &p.Price, &p.PriceType); err != nil {
			return nil, err
		}
		p.ID = &id
		prices = append(prices, p)
	}
	return prices, rows.Err()
}
