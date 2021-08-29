package main

import (
	"log"

	"github.com/MarceloZardoBR/go-api-frame/database"
	"github.com/MarceloZardoBR/go-api-frame/domain/service"
	"github.com/MarceloZardoBR/go-api-frame/infra/config"
	"github.com/MarceloZardoBR/go-api-frame/router"
	"github.com/MarceloZardoBR/go-api-frame/router/mainrouter/authorization"
	"github.com/gofiber/fiber/v2"

	"github.com/MarceloZardoBR/go-api-frame/server"
)

func main() {

	config, err := config.ReadAndLoadEnvVars()
	if err != nil {
		log.Println(err)
	}

	db, err := database.Instance()
	if err != nil {
		log.Println(err)
	}

	svc := service.NewServices(db, *config)

	userService := service.NewUserService(svc)

	_, err = userService.GetAll()
	if err != nil {
		log.Println(err)
	}

	authController := authorization.NewController()

	fiber := fiber.New()
	server := server.NewServer(fiber)

	router.AddRouter(fiber, authController)

	server.StartServer("5000")
}
