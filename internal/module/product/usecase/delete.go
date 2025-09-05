package usecase

import (
	"context"
	"fmt"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/google/uuid"
)

func (u *ProductUseCase) Delete(ctx context.Context, product model.Product) error {
	if product.ID == uuid.Nil {
		return helper.BadRequest("Id required")
	}

	result := u.db.Where("id = ?", product.ID).
		Where("seller_id = ?", product.SellerID).
		Delete(&entity.Product{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return helper.NotFound(fmt.Sprintf("Product not found"))
	}

	return nil
}
