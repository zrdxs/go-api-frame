package main

import (
	"log"

	"github.com/MarceloZardoBR/go-api-frame/domain/service"
	"github.com/MarceloZardoBR/go-api-frame/router"

	"github.com/gofiber/fiber/v2"

	"github.com/MarceloZardoBR/go-api-frame/database"
	"github.com/MarceloZardoBR/go-api-frame/server"
)

func main() {

	db, err := database.Instance()
	if err != nil {
		log.Println(err)
	}

	_ = service.NewServices(db)

	fiber := fiber.New()
	server := server.NewServer(fiber)

	router.AddRouter(fiber)

	server.StartServer("5000")
}
