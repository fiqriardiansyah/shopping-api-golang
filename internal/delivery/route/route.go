package route

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/middleware"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App
	*product.ProductController
	*order.OrderController
	*middleware.Middleware
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api")
	v1 := api.Group("/v1")

	c.SetupProductRoute(v1)
	c.SetupOrderRoute(v1)
}

func (c *RouteConfig) SetupProductRoute(router fiber.Router) {
	r := router.Group("/products", c.Middleware.AuthMiddleware)
	r.Get("", c.Middleware.RoleMiddleware("seller"), c.ProductController.FindAll)
	r.Get("/:id", c.Middleware.RoleMiddleware("seller"), c.ProductController.Find)
	r.Post("", c.Middleware.RoleMiddleware("seller"), c.ProductController.Create)
	r.Put("", c.Middleware.RoleMiddleware("seller"), c.ProductController.Update)
	r.Delete("/:id", c.Middleware.RoleMiddleware("seller"), c.ProductController.Delete)
}

func (c *RouteConfig) SetupOrderRoute(router fiber.Router) {
	r := router.Group("/orders", c.Middleware.AuthMiddleware)
	r.Post("/", c.Middleware.RoleMiddleware("buyer"), c.OrderController.MakeOrder)
	r.Get("/", c.Middleware.RoleMiddleware("buyer", "seller"), c.OrderController.MyOrder)
	r.Get("/seller", c.Middleware.RoleMiddleware("seller"), c.OrderController.MyOrderSeller)
}
