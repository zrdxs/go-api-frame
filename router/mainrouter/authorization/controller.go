package authorization

import (
	"fmt"

	"github.com/MarceloZardoBR/go-api-frame/infra/auth"
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Authorization(ctx *fiber.Ctx) error {
	auth := auth.AuthToken{}

	token, err := auth.GenerateToken()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Token: ", token)
	return ctx.SendString(token)
}
