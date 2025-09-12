package usecase

import (
	"context"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (u *OrderUseCase) MyOrder(ctx context.Context, userId uuid.UUID, query model.MyOrderRequest) (*[]model.MyOrderResponse, error) {
	db := u.db.Where("buyer_id = ?", userId)

	orders := []entity.Order{}

	if query.Status != "" {
		db = db.Where("Status = ?", query.Status)
	}

	if err := db.Preload("OrderItems.Product").Find(&orders).Error; err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	orderResponse := []model.MyOrderResponse{}
	if err := copier.Copy(&orderResponse, &orders); err != nil {
		return nil, helper.Internal(err.Error())
	}

	return &orderResponse, nil
}
