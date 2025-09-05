package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	SellerID    uuid.UUID  `gorm:"type:uuid;not null" json:"seller_id"`
	Name        string     `gorm:"type:varchar(200);not null" json:"name"`
	Description *string    `gorm:"type:text" json:"description"`
	Price       float64    `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int        `gorm:"default:0" json:"stock"`
	CategoryID  *uuid.UUID `gorm:"type:uuid" json:"category_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Foreign key relationships
	Category *Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnDelete:SET NULL" json:"category,omitempty"`

	// Relationships
	OrderItems []OrderItem `gorm:"foreignKey:ProductID" json:"order_items,omitempty"`
	CartItems  []CartItem  `gorm:"foreignKey:ProductID" json:"cart_items,omitempty"`
	Reviews    []Review    `gorm:"foreignKey:ProductID" json:"reviews,omitempty"`
}
