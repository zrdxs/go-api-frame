package main

import (
	"log"

	"github.com/MarceloZardoBR/go-api-frame/server"
	"github.com/MarceloZardoBR/go-api-frame/server/router"
	"github.com/MarceloZardoBR/go-api-frame/server/router/homeroute"
	"github.com/gofiber/fiber/v2"
)

func main() {

	server := initializeServer()
	server.Start()
}

// Initalize Echo, services and db configs
func initializeServer() *server.Server {

	fiber := fiber.New()

	homeController := homeroute.NewController(fiber)

	router.RegisterRoutes(fiber, homeController)

	server := server.Instance(fiber)

	log.Printf("Server running at port %d \n", 3000)

	return server
}
