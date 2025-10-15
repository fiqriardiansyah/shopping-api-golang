package middleware

import (
	"os"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/constant"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) RefreshMiddleware(ctx *fiber.Ctx) error {
	path := ctx.Path()

	accessToken := ctx.Cookies(constant.ACCESS_TOKEN)
	refreshToken := ctx.Cookies(constant.REFRESH_TOKEN)

	if accessToken == "" && refreshToken == "" {
		return ctx.Redirect(m.config.UserServiceUrl + "?redirect_uri=" + m.config.BaseUrl + path)
	}

	if _, e := helper.ValidateToken(accessToken, os.Getenv("JWT_SECRET")); e == nil {
		return ctx.Next()
	}

	if _, e := helper.ValidateToken(refreshToken, os.Getenv("JWT_REFRESH_TOKEN_SECRET")); e == nil {
		return ctx.Redirect(m.config.UserServiceUrl + "/refresh?redirect_uri=" + m.config.BaseUrl + path)
	}

	return ctx.Redirect(m.config.UserServiceUrl + "/logout?redirect_uri=" + m.config.BaseUrl + path)
}
