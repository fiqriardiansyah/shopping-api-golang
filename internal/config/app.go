package config

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/middleware"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/route"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *gorm.DB
	App       *fiber.App
	Validator *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	productController := product.InitializeProductHandler(config.DB, config.Validator)
	orderController := order.InitializeOrderHandler(config.DB, config.Validator)

	Route := route.RouteConfig{
		App:               config.App,
		ProductController: productController,
		OrderController:   orderController,
		Middleware:        middleware.NewMiddleware(config.DB),
	}

	Route.Setup()
}
