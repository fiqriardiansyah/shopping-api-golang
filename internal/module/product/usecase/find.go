package usecase

import (
	"context"
	"errors"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func (u *ProductUseCase) Find(ctx context.Context, product model.Product) (*model.Product, error) {
	entityProduct := entity.Product{}

	result := u.db.Where("id = ?", product.ID).Where("seller_id = ?", product.SellerID).First(&entityProduct)

	if result.Error != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	if err := copier.Copy(&product, &entityProduct); err != nil {
		return nil, err
	}

	return &product, nil
}
