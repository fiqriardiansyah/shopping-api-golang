package usecase

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product/repository"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	*repository.ProductRepository
	db *gorm.DB
}

func NewProductUseCase(productRepository *repository.ProductRepository, db *gorm.DB) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: productRepository,
		db:                db,
	}
}
