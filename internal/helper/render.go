package helper

import (
	"os"

	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
)

func Render(ctx *fiber.Ctx, view string, data fiber.Map) error {
	logoutHref := os.Getenv("USER_SERVICE_URL") + "/logout?redirect_uri=" + os.Getenv("BASE_URL")

	var user *model.User

	if u := ctx.Locals("user"); u != nil {
		if casted, ok := u.(*model.User); ok {
			user = casted
		}
	}

	globals := fiber.Map{
		"User":       user,
		"LogoutHref": logoutHref,
	}

	// merge page data with globals
	for k, v := range data {
		globals[k] = v
	}

	return ctx.Render(view, globals, "layout/base")
}
