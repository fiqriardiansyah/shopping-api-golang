package usecase

import (
	"context"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/jinzhu/copier"
)

func (u *ProductUseCase) PageCategories(ctx context.Context) ([]model.Category, error) {
	entityCategory := []entity.Category{}
	categoy := []model.Category{}
	db := u.db

	result := db.Find(&entityCategory)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := copier.Copy(&categoy, &entityCategory); err != nil {
		return nil, err
	}

	return categoy, nil
}
