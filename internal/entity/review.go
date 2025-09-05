package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BuyerID   uuid.UUID `gorm:"type:uuid;not null" json:"buyer_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Rating    int       `gorm:"not null;check:rating BETWEEN 1 AND 5" json:"rating"`
	Comment   *string   `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `json:"created_at"`

	// Foreign key relationships
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
}
