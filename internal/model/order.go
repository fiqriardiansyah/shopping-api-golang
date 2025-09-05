package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID          uuid.UUID `json:"id"`
	BuyerID     uuid.UUID `json:"-"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	OrderItems []OrderItem `json:"order_items,omitempty"`
}

type MyOrderRequest struct {
	Status string `json:"status"`
}

type MyOrderSellerRequest struct {
	Status string `json:"status"`
}
