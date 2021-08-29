package router

import (
	"github.com/MarceloZardoBR/go-api-frame/router/mainrouter/authorization"
	"github.com/gofiber/fiber/v2"
)

func AddRouter(app *fiber.App, authController *authorization.Controller) {

	rootGroup := app.Group("")

	rootGroup.Get("/health/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("OK")
	})

	mainGroup := rootGroup.Group("main")

	mainGroup.Get("/authorization/", authController.Authorization)
}
