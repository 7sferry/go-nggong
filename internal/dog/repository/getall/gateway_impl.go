package getall

import (
	"context"
	"database/sql"

	"go-nggong/internal/dog/domain"
)

type GatewayImpl struct {
	DB *sql.DB
}

func New(db *sql.DB) *GatewayImpl {
	return &GatewayImpl{DB: db}
}

func (g *GatewayImpl) FindAll(ctx context.Context) ([]domain.DogRecord, error) {
	rows, err := g.DB.QueryContext(ctx, "SELECT id, name, email, role FROM dogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dogs []domain.DogRecord
	for rows.Next() {
		var id int
		var d domain.DogRecord
		if err := rows.Scan(&id, &d.Name, &d.Email, &d.Role); err != nil {
			return nil, err
		}
		d.ID = &id
		dogs = append(dogs, d)
	}
	return dogs, rows.Err()
}
