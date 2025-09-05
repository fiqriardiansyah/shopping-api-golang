package model

import "github.com/google/uuid"

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`

	// Foreign key relationships
	Cart    Cart    `gorm:"foreignKey:CartID;references:ID;constraint:OnDelete:CASCADE" json:"cart,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
}
