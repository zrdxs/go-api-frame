package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Server holds server struct
type Server struct {
	fiber *fiber.App
}

// NewServer returns new server instance
func NewServer(fiber *fiber.App) *Server {
	return &Server{
		fiber: fiber,
	}
}

// StartServer configure fiber HTTP
func (s *Server) StartServer(port string) error {

	err := s.fiber.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	log.Println("Starting Server at port: %d", port)

	return nil
}
