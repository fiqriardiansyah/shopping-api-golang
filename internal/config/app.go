package config

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/http/route"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/user"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/ui/page"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	App        *fiber.App
	Validator  *validator.Validate
	Config     *helper.Config
	GrpcClient *helper.GrpcClient
}

func Bootstrap(config *BootstrapConfig) {
	productController := product.InitializeProductHandler(config.DB, config.Validator)
	orderController := order.InitializeOrderHandler(config.DB, config.Validator)
	userController := user.InitializeUserHandler(config.GrpcClient)

	page := page.NewPages(orderController, productController, userController, config.Config)

	Route := route.RouteConfig{
		App:        config.App,
		Product:    productController,
		Order:      orderController,
		User:       userController,
		Middleware: middleware.NewMiddleware(config.DB, config.Config),
		Page:       page,
	}

	Route.Setup()
}
