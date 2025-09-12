package page

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (p *Pages) PageProducts(ctx *fiber.Ctx) error {
	query := ctx.Query("query")
	category := ctx.Query("category")

	products, err := p.Product.ProductUseCase.PageProductList(ctx.UserContext(), query, category)

	if err != nil {
		return helper.Render(ctx, "page/products", fiber.Map{
			"Message": err.Error(),
		})
	}

	categories, err := p.Product.ProductUseCase.PageCategories(ctx.UserContext())

	if err != nil {
		return helper.Render(ctx, "page/products", fiber.Map{
			"Message": err.Error(),
		})
	}

	return helper.Render(ctx, "page/products", fiber.Map{
		"Products":   products,
		"Categories": categories,
		"Query":      query,
		"Category":   category,
	})
}
