package middleware

import (
	"os"
	"strings"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AuthMiddleware(ctx *fiber.Ctx) error {

	authorization := ctx.Get("Authorization")
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		return helper.Unauthorized("Authorization token not valid or not found")
	}
	token := strings.TrimPrefix(authorization, "Bearer ")

	claims, err := helper.ValidateToken(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	user := model.User{
		Id:    claims.UserId,
		Email: claims.Email,
		Roles: claims.Roles,
	}

	ctx.Locals("user", &user)

	return ctx.Next()
}
