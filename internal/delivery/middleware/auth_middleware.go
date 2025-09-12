package middleware

import (
	"os"
	"strings"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/constant"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AuthMiddleware(ctx *fiber.Ctx) error {

	var token string

	accessTokenCookie := ctx.Cookies(constant.ACCESS_TOKEN)

	authorization := ctx.Get("Authorization")

	if accessTokenCookie != "" {
		token = accessTokenCookie
	} else if authorization != "" {
		if !strings.HasPrefix(authorization, "Bearer ") {
			return helper.Unauthorized("Authorization token not valid or not found")
		}
		token = strings.TrimPrefix(authorization, "Bearer ")
	} else {
		return helper.Unauthorized("Authorization token not valid or not found")
	}

	claims, err := helper.ValidateToken(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		return helper.Unauthorized(err.Error())
	}

	user := model.User{
		Id:    claims.UserId,
		Email: claims.Email,
		Roles: claims.Roles,
	}

	ctx.Locals("user", &user)

	return ctx.Next()
}
