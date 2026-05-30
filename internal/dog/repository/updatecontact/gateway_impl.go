package updatecontact

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

func (g *GatewayImpl) UpdateByDogID(ctx context.Context, contact domain.ContactRecord) (*domain.ContactRecord, error) {
	var id int
	err := g.DB.QueryRowContext(ctx,
		"UPDATE dog_contacts SET phone = $1, address = $2, emergency = $3 WHERE dog_id = $4 RETURNING id",
		contact.Phone, contact.Address, contact.Emergency, contact.DogID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	contact.ID = &id
	return &contact, nil
}
