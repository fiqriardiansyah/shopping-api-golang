package page

import (
	"os"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/constant"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (p *Pages) PageHome(ctx *fiber.Ctx) error {
	loginHref := p.Config.UserServiceUrl + "?redirect_uri=" + p.Config.BaseUrl
	registerHref := p.Config.UserServiceUrl + "/register?redirect_uri=" + p.Config.BaseUrl
	logoutHref := p.Config.UserServiceUrl + "/logout?redirect_uri=" + p.Config.BaseUrl

	if claims, e := helper.ValidateToken(ctx.Cookies(constant.ACCESS_TOKEN), os.Getenv("JWT_SECRET")); e == nil {
		// user, errUser := p.User.UseCase.GetUser(claims.UserId)

		// if errUser != nil {

		// }

		return ctx.Render("page/index", fiber.Map{
			"LoginHref":    loginHref,
			"RegisterHref": registerHref,
			"LogoutHref":   logoutHref,
			"User":         fiber.Map{"Email": claims.Email},
		})
	}

	return ctx.Render("page/index", fiber.Map{
		"LoginHref":    loginHref,
		"RegisterHref": registerHref,
	})
}
