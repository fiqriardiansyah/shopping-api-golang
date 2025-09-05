package usecase

import (
	"context"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/google/uuid"
	"time"
)

type MyOrderSellerResponse struct {
	ProductName string    `json:"product_name"`
	ProductId   uuid.UUID `json:"product_id"`
	BuyerId     uuid.UUID `json:"buyer_id"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	Quantity    float64   `json:"quantity"`
	PriceItem   float64   `json:"price_item"`
}

func (u *OrderUseCase) MyOrderSeller(ctx context.Context, userId uuid.UUID, query model.MyOrderSellerRequest) (*[]MyOrderSellerResponse, error) {
	rows := []MyOrderSellerResponse{}

	builder := u.db.Table("order_items as oi").
		Select("p.name as product_name, p.id as product_id, o.buyer_id as buyer_id, o.status, o.total_amount, o.created_at, oi.quantity, oi.price as price_item").
		Joins("JOIN orders as o on o.id = oi.order_id").
		Joins("JOIN products as p on oi.product_id = p.id").
		Where("p.seller_id = ?", userId)

	if query.Status != "" {
		builder = builder.Where("o.status = ?", query.Status)
	}

	if err := builder.Scan(&rows).Error; err != nil {
		return nil, err
	}

	return &rows, nil
}
