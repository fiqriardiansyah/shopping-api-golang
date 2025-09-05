package entity

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BuyerID   uuid.UUID `gorm:"type:uuid;not null;unique" json:"buyer_id"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cart_items,omitempty"`
}
