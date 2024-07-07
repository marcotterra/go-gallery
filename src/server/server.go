package server

import (
	"go-gallery/src/core/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	folderHandlers ports.FolderHandler
}

func NewServer(fHandler ports.FolderHandler) *Server {
	return &Server{
		folderHandlers: fHandler,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"hello": "world",
		})
	})

	v1 := app.Group("/v1")

	folderRoutes := v1.Group("/user")
	folderRoutes.Post("/", s.folderHandlers.Store)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
