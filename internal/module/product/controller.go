package product

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/module/product/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	*usecase.ProductUseCase
	*validator.Validate
}

func NewProductController(useCase *usecase.ProductUseCase, validate *validator.Validate) *ProductController {
	return &ProductController{
		ProductUseCase: useCase,
		Validate:       validate,
	}
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	product := model.Product{
		SellerID: user.Id,
	}

	if err := ctx.BodyParser(&product); err != nil {
		return helper.Internal(err.Error())
	}

	id, err := c.ProductUseCase.Create(ctx.UserContext(), product)
	if err != nil {
		return helper.Internal(err.Error())
	}

	return helper.Success(ctx, fiber.Map{
		"id": id,
	}, 200)
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	product := model.Product{
		SellerID: user.Id,
	}

	if err := ctx.BodyParser(&product); err != nil {
		return helper.Internal(err.Error())
	}

	id, err := c.ProductUseCase.Update(ctx.UserContext(), product)
	if err != nil {
		return helper.Internal(err.Error())
	}

	return helper.Success(ctx, fiber.Map{
		"id": id,
	}, 200)
}

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	product := model.Product{
		SellerID: user.Id,
	}

	id, err := uuid.Parse(ctx.Params("id"))
	product.ID = id

	err = c.ProductUseCase.Delete(ctx.UserContext(), product)
	if err != nil {
		return helper.Internal(err.Error())
	}

	return helper.Success(ctx, fiber.Map{
		"id": product.ID,
	}, 200)
}

func (c *ProductController) Find(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	product := model.Product{
		SellerID: user.Id,
	}

	id, err := uuid.Parse(ctx.Params("id"))
	product.ID = id

	result, err := c.ProductUseCase.Find(ctx.UserContext(), product)
	if err != nil {
		return helper.Internal(err.Error())
	}

	return helper.Success(ctx, &result, 200)
}

func (c *ProductController) FindAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)

	param := model.ProductParam{}
	if err := ctx.QueryParser(&param); err != nil {
		return helper.BadRequest(err.Error())
	}

	result, err := c.ProductUseCase.FindAll(ctx.UserContext(), user.Id, param)
	if err != nil {
		return helper.Internal(err.Error())
	}

	return helper.Success(ctx, &result, 200)
}
