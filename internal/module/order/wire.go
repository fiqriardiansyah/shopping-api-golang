//go:build wireinject
// +build wireinject

package order

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order/repository"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeOrderHandler(db *gorm.DB, validate *validator.Validate) *OrderController {
	wire.Build(NewOrderController, usecase.NewOrderUseCase, repository.NewOrderRepository)
	return nil
}
