package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	SellerID    uuid.UUID `json:"-"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	Stock       int       `json:"stock" validate:"required"`
	CategoryID  uuid.UUID `json:"category_id" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`

	// Relationships
	OrderItems []OrderItem `json:"order_items,omitempty"`
	CartItems  []CartItem  `json:"cart_items,omitempty"`
	Reviews    []Review    `json:"reviews,omitempty"`
}

type ProductParam struct {
	Query string `json:"query"`
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CategoryID  uuid.UUID `json:"category_id" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
