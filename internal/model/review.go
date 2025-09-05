package model

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID `json:"id"`
	BuyerID   uuid.UUID `json:"buyer_id"`
	ProductID uuid.UUID `json:"product_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
