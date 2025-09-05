package usecase

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order/repository"
	"gorm.io/gorm"
)

type OrderUseCase struct {
	*repository.OrderRepository
	db *gorm.DB
}

func NewOrderUseCase(repo *repository.OrderRepository, db *gorm.DB) *OrderUseCase {
	return &OrderUseCase{
		OrderRepository: repo,
		db:              db,
	}
}
