//go:build wireinject
// +build wireinject

package product

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product/repository"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeProductHandler(db *gorm.DB, validate *validator.Validate) *ProductController {
	wire.Build(repository.NewProductRepository, usecase.NewProductUseCase, NewProductController)
	return nil
}
