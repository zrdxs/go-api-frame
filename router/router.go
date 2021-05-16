package router

import (
	"github.com/gofiber/fiber/v2"
)

func AddRouter(app *fiber.App) {

	rootGroup := app.Group("")

	rootGroup.Get("/health/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("OK")
	})
}
