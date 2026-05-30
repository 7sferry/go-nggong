package getcontact

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

func (g *GatewayImpl) FindByDogID(ctx context.Context, dogID int) (*domain.ContactRecord, error) {
	var id int
	var c domain.ContactRecord
	err := g.DB.QueryRowContext(ctx,
		"SELECT id, dog_id, phone, address, emergency FROM dog_contacts WHERE dog_id = $1", dogID).
		Scan(&id, &c.DogID, &c.Phone, &c.Address, &c.Emergency)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	c.ID = &id
	return &c, nil
}
