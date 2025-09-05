package usecase

import (
	"context"
	"fmt"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

func (u *OrderUseCase) MakeOrder(ctx context.Context, req []model.OrderRequest) (string, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var errs []string
	var total float64

	orderItems := []entity.OrderItem{}

	for _, item := range req {
		findProduct := entity.Product{}
		if err := tx.Where("id = ?", item.ProductID).First(&findProduct).Error; err != nil {
			errs = append(errs, fmt.Sprintf("product id %v does not exist", item.ProductID))
			continue
		}

		//can't buy their own product
		if findProduct.SellerID == item.BuyerID {
			errs = append(errs, fmt.Sprintf("Can't make order on item id %s", item.ProductID))
			continue
		}

		if item.Quantity > findProduct.Stock {
			errs = append(errs, fmt.Sprintf("Stock product not enough on item id"))
			continue
		}

		orderItems = append(orderItems, entity.OrderItem{
			Quantity:  item.Quantity,
			ProductID: findProduct.ID,
			Price:     findProduct.Price,
		})

		total = total + float64(item.Quantity)*findProduct.Price
	}

	if len(errs) > 0 {
		return "", helper.BadRequest(strings.Join(errs, ", "))
	}

	errs = []string{}

	for _, item := range orderItems {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", item.ProductID).
			First(&entity.Product{}).
			Error; err != nil {
			errs = append(errs, err.Error())
			continue
		}

		if err := tx.Model(&entity.Product{}).Where("id = ?", item.ProductID).
			Update("Stock", gorm.Expr("Stock - ?", item.Quantity)).Error; err != nil {
			errs = append(errs, err.Error())
			continue
		}
	}

	if len(errs) > 0 {
		return "", helper.Internal(strings.Join(errs, ", "))
	}

	order := entity.Order{
		BuyerID:     req[0].BuyerID,
		OrderItems:  orderItems,
		TotalAmount: total,
	}

	if err := tx.Create(&order).Error; err != nil {
		return "", helper.Internal(err.Error())
	}

	tx.Commit()

	return order.ID.String(), nil
}
