package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func (u *ProductUseCase) Update(ctx context.Context, product model.Product) (string, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if product.ID == uuid.Nil {
		return "", helper.BadRequest("Product id not found")
	}

	entityProduct := entity.Product{}

	result := tx.Where("id = ?", product.ID).Where("seller_id = ?", product.SellerID).First(&entityProduct)
	if result.Error != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", helper.Internal(fmt.Sprintf("Product with id %s not found", product.ID))
	}

	if err := copier.CopyWithOption(&entityProduct, &product, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return "", err
	}

	if err := u.ProductRepository.Update(tx, &entityProduct); err != nil {
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return entityProduct.ID.String(), nil
}
