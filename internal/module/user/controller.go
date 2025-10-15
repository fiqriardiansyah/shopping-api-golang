package user

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/user/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	GrpcClient *helper.GrpcClient
	UseCase    *usecase.UserUseCase
}

func NewUserController(client *helper.GrpcClient, u *usecase.UserUseCase) *UserController {
	return &UserController{
		GrpcClient: client,
		UseCase:    u,
	}
}

func (c *UserController) RegisterRoutes(router fiber.Router, mw *middleware.Middleware) {
	r := router.Group("/user")
	r.Get("/:id", c.GetUser)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	userResponse, err := c.UseCase.GetUser(user.Id)
	if err != nil {
		return ctx.JSON(err.Error())
	}
	return ctx.JSON(userResponse)
}
