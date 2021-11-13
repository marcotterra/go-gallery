package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Gallery struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

type Galleries struct {
	Galleries []Gallery `json:"galleries"`
}

func main() {
	app := fiber.New()

	store := Galleries{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"hello": "world",
		})
	})

	app.Get("/galleries", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": store.Galleries,
		})
	})

	app.Post("/galleries", func(c *fiber.Ctx) error {
		gallery := new(Gallery)

		if err := c.BodyParser(gallery); err != nil {
			log.Println(err)
			return c.Status(400).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

		store.Galleries = append(store.Galleries, Gallery{ID: gallery.ID, Description: gallery.Description})

		return c.JSON(&fiber.Map{
			"data": store.Galleries,
		})
	})

	app.Put("/galleries/:id", func(c *fiber.Ctx) error {
		gallery := new(Gallery)

		if err := c.BodyParser(gallery); err != nil {
			log.Println(err)
			return c.Status(400).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

	})

	log.Fatal(app.Listen(":8000"))
}
