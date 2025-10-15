package middleware

import (
	"slices"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) RoleMiddleware(roles ...string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*model.User)

		for _, role := range user.Roles {
			if slices.Contains(roles, role) {
				return ctx.Next()
			}
		}

		return helper.Forbidden("Forbidden resource")
	}
}
