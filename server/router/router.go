package router

import (
	"github.com/MarceloZardoBR/go-api-frame/server/router/homeroute"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes register routes with the controllers handlers
func RegisterRoutes(fiber *fiber.App, homeController *homeroute.Controller) {
	fibergroup := fiber.Group("")

	fibergroup.Get("/home/", homeController.HandleGetHome)
}
