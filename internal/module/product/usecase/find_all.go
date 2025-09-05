package usecase

import (
	"context"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (u *ProductUseCase) FindAll(ctx context.Context, sellerId uuid.UUID, param model.ProductParam) ([]model.Product, error) {
	entityProducts := []entity.Product{}
	products := []model.Product{}

	db := u.db.Where("seller_id = ?", sellerId)

	if param.Query != "" {
		db = db.Where("name ILIKE ? OR description ILIKE ?", "%"+param.Query+"%", "%"+param.Query+"%")
	}

	result := db.Find(&entityProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := copier.Copy(&products, &entityProducts); err != nil {
		return nil, err
	}

	return products, nil
}
