package getbyid

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

func (g *GatewayImpl) FindByID(ctx context.Context, id int) (*domain.DogRecord, error) {
	var rid int
	var d domain.DogRecord
	err := g.DB.QueryRowContext(ctx,
		"SELECT id, name, email, role FROM dogs WHERE id = $1", id).
		Scan(&rid, &d.Name, &d.Email, &d.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	d.ID = &rid
	return &d, nil
}
