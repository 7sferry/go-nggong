package update

import (
	"context"
	"database/sql"
	"errors"

	"go-nggong/internal/dog/domain"
)

type GatewayImpl struct {
	DB *sql.DB
}

func New(db *sql.DB) *GatewayImpl {
	return &GatewayImpl{DB: db}
}

func (g *GatewayImpl) Update(ctx context.Context, dog domain.DogRecord) (*domain.DogRecord, error) {
	if dog.ID == nil {
		return nil, errors.New("id required")
	}
	res, err := g.DB.ExecContext(ctx,
		"UPDATE dogs SET name = $1, email = $2, role = $3 WHERE id = $4",
		dog.Name, dog.Email, dog.Role, *dog.ID)
	if err != nil {
		return nil, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	return &dog, nil
}
