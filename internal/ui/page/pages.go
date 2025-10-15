package page

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/user"
	"github.com/gofiber/fiber/v2"
)

type Pages struct {
	Order   *order.OrderController
	Product *product.ProductController
	User    *user.UserController
	Config  *helper.Config
}

func NewPages(order *order.OrderController, product *product.ProductController, user *user.UserController, config *helper.Config) *Pages {
	return &Pages{
		Order:   order,
		Product: product,
		User:    user,
		Config:  config,
	}
}

func (p *Pages) RegisterPages(router fiber.Router, mw *middleware.Middleware) {

	protected := router.Group("/", mw.RefreshMiddleware, mw.AuthMiddleware)

	protected.Get("/", p.PageHome)

	productRoute := protected.Group("/products")
	productRoute.Get("/", mw.RoleMiddleware("buyer"), p.PageProducts)
}
