package repository

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
)

type ProductRepository struct {
	helper.Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
