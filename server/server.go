package server

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

const (
	port = 3000
)

var (
	once     sync.Once
	instance *Server
)

// Server holds server struct
type Server struct {
	Fiber *fiber.App
}

// Instance start a new Server instance
func Instance(fiber *fiber.App) *Server {
	once.Do(func() {
		instance = &Server{
			Fiber: fiber,
		}
	})
	return instance
}

// Start server at defined port
func (s *Server) Start() error {

	err := s.Fiber.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	return nil
}
