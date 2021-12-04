package ports

import "github.com/gofiber/fiber/v2"

type FolderService interface {
	Store(title string) error
}

type FolderRepository interface {
	Store(title string) error
}

type FolderHandler interface {
	Store(c *fiber.Ctx) error
}
