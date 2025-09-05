package entity

import "github.com/google/uuid"

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity  int       `gorm:"not null;check:quantity > 0" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`

	// Foreign key relationships
	Order   Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE" json:"order,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
}
