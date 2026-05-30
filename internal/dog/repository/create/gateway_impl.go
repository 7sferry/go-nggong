package create

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

func (g *GatewayImpl) Save(ctx context.Context, dog domain.DogRecord) (domain.DogRecord, error) {
	var id int
	err := g.DB.QueryRowContext(ctx,
		"INSERT INTO dogs (name, email, role) VALUES ($1, $2, $3) RETURNING id",
		dog.Name, dog.Email, dog.Role).Scan(&id)
	if err != nil {
		return domain.DogRecord{}, err
	}
	dog.ID = &id
	return dog, nil
}
