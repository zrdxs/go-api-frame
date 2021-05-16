package main

import (
	"log"

	"github.com/MarceloZardoBR/go-api-frame/router"

	"github.com/gofiber/fiber/v2"

	"github.com/MarceloZardoBR/go-api-frame/database"
	"github.com/MarceloZardoBR/go-api-frame/server"
)

func main() {

	psqlconn := database.CreateConfiguration("localhost", 5432, "postgres", "", "db_teste")
	_, err := database.StartDB(psqlconn)
	if err != nil {
		log.Println(err)
	}

	fiber := fiber.New()
	server := server.NewServer(fiber)

	router.AddRouter(fiber)

	server.StartServer("5000")
}
