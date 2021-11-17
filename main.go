package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Gallery struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

type Store struct {
	amount    int       `json:"amount,omitempty"`
	Galleries []Gallery `json:"galleries"`
}

// Append an item in slice
func (g *Store) Create(desc string) {
	g.amount++
	g.Galleries = append(g.Galleries, Gallery{ID: g.amount, Description: desc})
}

// Get an item from slice
func (g *Store) Read(id int) (updated Gallery, ok bool) {
	for k := range g.Galleries {
		if g.Galleries[k].ID == id {
			return g.Galleries[k], true
		}
	}

	return Gallery{}, false
}

// Update an item in slice
func (g *Store) Update(gallery Gallery) (updated Gallery, ok bool) {
	for k := range g.Galleries {
		if g.Galleries[k].ID == gallery.ID {
			g.Galleries[k] = gallery

			return g.Galleries[k], true
		}
	}

	return Gallery{}, false
}

// Delete an item in slice
func (g *Store) Delete(id int) (ok bool) {
	for k := range g.Galleries {
		if g.Galleries[k].ID == id {
			// g.amount--
			g.Galleries = append(g.Galleries[:k], g.Galleries[k+1:]...)

			return true
		}
	}

	return false
}

func main() {
	app := fiber.New()

	store := Store{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"hello": "world",
		})
	})

	app.Get("/galleries", func(c *fiber.Ctx) error {
		return c.JSON(store)
	})

	app.Get("/galleries/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		if id == "" {
			return c.Status(400).JSON(&fiber.Map{
				"message": "id wasnt provided",
			})
		}

		iid, _ := strconv.ParseInt(id, 8, 32)

		found, ok := store.Read(int(iid))

		if !ok {
			return c.Status(400).JSON(&fiber.Map{
				"message": "id not found",
			})
		}

		return c.JSON(found)
	})

	app.Post("/galleries", func(c *fiber.Ctx) error {
		gallery := &Gallery{}

		if err := c.BodyParser(gallery); err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

		store.Create(gallery.Description)

		return c.JSON(gallery)
	})

	app.Put("/galleries/:id", func(c *fiber.Ctx) error {
		gallery := &Gallery{}
		id := c.Params("id")

		if err := c.BodyParser(gallery); err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

		if id == "" {
			return c.Status(400).JSON(&fiber.Map{
				"message": "id wasnt provided",
			})
		}

		intId, _ := strconv.ParseInt(id, 8, 32)

		updated, ok := store.Update(Gallery{ID: int(intId), Description: gallery.Description})

		if !ok {
			return c.Status(400).JSON(&fiber.Map{
				"message": fmt.Sprintf("%d not found", gallery.ID),
			})
		}

		return c.JSON(updated)
	})

	app.Delete("/galleries/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		if id == "" {
			return c.Status(400).JSON(&fiber.Map{
				"message": "id wasnt provided",
			})
		}

		intId, _ := strconv.ParseInt(id, 8, 32)

		if ok := store.Delete(int(intId)); !ok {
			return c.Status(400).JSON(&fiber.Map{
				"message": fmt.Sprintf("%d not found", intId),
			})
		}

		return c.Status(204).JSON(&fiber.Map{
			"message": fmt.Sprintf("%d was removed", intId),
		})
	})

	log.Fatal(app.Listen(":8000"))
}
