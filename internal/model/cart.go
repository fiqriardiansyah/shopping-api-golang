package model

import (
	"github.com/google/uuid"
	"time"
)

type Cart struct {
	ID        uuid.UUID `json:"id"`
	BuyerID   uuid.UUID `json:"buyer_id"`
	CreatedAt time.Time `json:"created_at"`

	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cart_items,omitempty"`
}
