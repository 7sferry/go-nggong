package createcontact

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

func (g *GatewayImpl) Save(ctx context.Context, contact domain.ContactRecord) (domain.ContactRecord, error) {
	var id int
	err := g.DB.QueryRowContext(ctx,
		"INSERT INTO dog_contacts (dog_id, phone, address, emergency) VALUES ($1, $2, $3, $4) RETURNING id",
		contact.DogID, contact.Phone, contact.Address, contact.Emergency).Scan(&id)
	if err != nil {
		return domain.ContactRecord{}, err
	}
	contact.ID = &id
	return contact, nil
}
