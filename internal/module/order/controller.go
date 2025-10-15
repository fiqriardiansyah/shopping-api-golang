package order

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/order/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OrderController struct {
	*usecase.OrderUseCase
}

func (c *OrderController) RegisterRoutes(router fiber.Router, mw *middleware.Middleware) {
	r := router.Group("/orders", mw.AuthMiddleware)
	r.Post("/", mw.RoleMiddleware("buyer"), c.MakeOrder)
	r.Get("/", mw.RoleMiddleware("buyer", "seller"), c.MyOrder)
	r.Get("/seller", mw.RoleMiddleware("seller"), c.MyOrderSeller)
}

func NewOrderController(useCase *usecase.OrderUseCase) *OrderController {
	return &OrderController{
		OrderUseCase: useCase,
	}
}

func (c *OrderController) MakeOrder(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	orderItems := []model.OrderRequest{}

	if err := ctx.BodyParser(&orderItems); err != nil {
		return helper.BadRequest(err.Error())
	}

	for i, _ := range orderItems {
		orderItems[i].BuyerID = user.Id
	}

	id, err := c.OrderUseCase.MakeOrder(ctx.UserContext(), orderItems)
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	return helper.Success(ctx, id, 200)
}

func (c *OrderController) MyOrder(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	query := model.MyOrderRequest{}

	if err := ctx.QueryParser(&query); err != nil {
		return helper.BadRequest(err.Error())
	}

	orders, err := c.OrderUseCase.MyOrder(ctx.UserContext(), user.Id, query)
	if err != nil {
		return err
	}

	return helper.Success(ctx, orders, 200)
}

func (c *OrderController) MyOrderSeller(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	query := model.MyOrderSellerRequest{}

	logrus.Info(user.Email)

	if err := ctx.QueryParser(&query); err != nil {
		return helper.BadRequest(err.Error())
	}

	orders, err := c.OrderUseCase.MyOrderSeller(ctx.UserContext(), user.Id, query)
	if err != nil {
		return err
	}

	return helper.Success(ctx, orders, 200)
}
