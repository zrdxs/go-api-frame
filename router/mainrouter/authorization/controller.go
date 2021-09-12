package authorization

import (
	"net/http"

	"github.com/MarceloZardoBR/go-api-frame/infra/auth"
	"github.com/MarceloZardoBR/go-api-frame/infra/config"
	"github.com/gofiber/fiber/v2"
)

// Controller holds entity struct
type Controller struct {
	cfg *config.Config
}

// NewController return new entity instance
func NewController(cfg *config.Config) *Controller {
	return &Controller{
		cfg: cfg,
	}
}

func (c *Controller) Authorization(ctx *fiber.Ctx) error {
	auth := auth.AuthToken{}

	authResponse, err := auth.GenerateUserTokenAndResponse(c.cfg)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return ctx.JSON(authResponse)
}
