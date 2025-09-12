package usecase

import (
	"context"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/jinzhu/copier"
)

func (u *ProductUseCase) PageProductList(ctx context.Context, query string, category string) ([]model.Product, error) {
	entityProducts := []entity.Product{}
	products := []model.Product{}

	db := u.db

	if query != "" {
		db = db.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if category != "" && category != "all" {
		db = db.Where("category_id = ?", category)
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
