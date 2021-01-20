package homeroute

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	once     sync.Once
	instance *Controller
)

// Controller for home route
type Controller struct {
	fiber *fiber.App
}

// NewController returns new controller instance
func NewController(fiber *fiber.App) *Controller {
	once.Do(func() {
		instance = &Controller{
			fiber: fiber,
		}
	})
	return instance
}

func (c *Controller) HandleGetHome(ctx *fiber.Ctx) error {

	return ctx.JSON("Route Home GET")
}
