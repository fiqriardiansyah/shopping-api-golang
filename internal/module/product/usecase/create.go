package usecase

import (
	"context"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/jinzhu/copier"
)

func (u *ProductUseCase) Create(ctx context.Context, product model.Product) (string, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	entityProduct := entity.Product{}

	if err := copier.Copy(&entityProduct, &product); err != nil {
		return "", err
	}

	if err := u.ProductRepository.Create(tx, &entityProduct); err != nil {
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return entityProduct.ID.String(), nil
}
