package usecase

import (
	"context"
	"time"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/google/uuid"
)

type MyOrderSellerResponse struct {
	ProductName     string    `json:"product_name"`
	ProductId       uuid.UUID `json:"product_id"`
	OrderId         uuid.UUID `json:"order_id"`
	BuyerId         uuid.UUID `json:"buyer_id"`
	Status          string    `json:"status"`
	TotalAmount     float64   `json:"total_amount"`
	CreatedAt       time.Time `json:"created_at"`
	ProductQuantity float64   `json:"product_quantity"`
	ProductPrice    float64   `json:"product_price"`
}

type Item struct {
	ProductName     string    `json:"product_name"`
	ProductId       uuid.UUID `json:"product_id"`
	ProductPrice    float64   `json:"product_price"`
	ProductQuantity float64   `json:"product_quantity"`
}

type OrdersResponse struct {
	OrderId     uuid.UUID `json:"order_id"`
	BuyerId     uuid.UUID `json:"buyer_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	TotalAmount float64   `json:"total_amount"`
	Items       []Item    `json:"items"`
}

func (u *OrderUseCase) MyOrderSeller(ctx context.Context, userId uuid.UUID, query model.MyOrderSellerRequest) (*[]OrdersResponse, error) {
	rows := []MyOrderSellerResponse{}

	builder := u.db.Table("order_items as oi").
		Select("o.id as order_id, p.name as product_name, p.id as product_id, o.buyer_id as buyer_id, o.status, o.total_amount, o.created_at, oi.quantity as product_quantity, oi.price as product_price").
		Joins("JOIN orders as o on o.id = oi.order_id").
		Joins("JOIN products as p on oi.product_id = p.id").
		Where("p.seller_id = ?", userId)

	if query.Status != "" {
		builder = builder.Where("o.status = ?", query.Status)
	}

	if err := builder.Scan(&rows).Error; err != nil {
		return nil, err
	}

	orders := []OrdersResponse{}

	for _, row := range rows {
		exist := false
		for i, o := range orders {
			if o.OrderId == row.OrderId {
				exist = true
				orders[i].Items = append(orders[i].Items, Item{
					ProductName:     row.ProductName,
					ProductId:       row.ProductId,
					ProductPrice:    row.ProductPrice,
					ProductQuantity: row.ProductQuantity,
				})
			}
		}
		if !exist {
			orders = append(orders, OrdersResponse{
				OrderId:     row.OrderId,
				BuyerId:     row.BuyerId,
				Status:      row.Status,
				CreatedAt:   row.CreatedAt,
				TotalAmount: row.TotalAmount,
				Items: []Item{
					{ProductName: row.ProductName,
						ProductId:       row.ProductId,
						ProductPrice:    row.ProductPrice,
						ProductQuantity: row.ProductQuantity,
					},
				},
			})
		}
	}

	return &orders, nil
}
