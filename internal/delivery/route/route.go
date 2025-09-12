package route

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/middleware"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/ui/page"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App        *fiber.App
	Product    *product.ProductController
	Order      *order.OrderController
	Middleware *middleware.Middleware
	Page       *page.Pages
}

func (c *RouteConfig) Setup() {
	// UI
	c.Page.RegisterPages(c.App, c.Middleware)

	// API
	api := c.App.Group("/api")
	v1 := api.Group("/v1")
	c.Product.RegisterRoutes(v1, c.Middleware)
	c.Order.RegisterRoutes(v1, c.Middleware)
}
