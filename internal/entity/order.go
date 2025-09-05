package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BuyerID     uuid.UUID `gorm:"type:uuid;not null" json:"buyer_id"`
	Status      string    `gorm:"type:varchar(20);check:status IN ('PENDING','PAID','SHIPPED','COMPLETED','CANCELLED');default:'PENDING'" json:"status"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}
