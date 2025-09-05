package entity

import "github.com/google/uuid"

type CartItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CartID    uuid.UUID `gorm:"type:uuid;not null" json:"cart_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null;check:quantity > 0" json:"quantity"`

	// Foreign key relationships
	Cart    Cart    `gorm:"foreignKey:CartID;references:ID;constraint:OnDelete:CASCADE" json:"cart,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
}
