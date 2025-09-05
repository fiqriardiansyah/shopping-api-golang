package config

import (
	"github.com/fiqriardiansyah/shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func NewFiber() (*fiber.App, error) {

	fiberPrefork, err := strconv.ParseBool(os.Getenv("FIBER_PREFORK"))
	if err != nil {
		return nil, err
	}

	var app = fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
		Prefork: fiberPrefork,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*helper.AppError); ok {
				return ctx.Status(e.Code).JSON(e)
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Status:  "INTERNAL SERVER ERROR",
			})
		},
	})

	return app, nil
}
