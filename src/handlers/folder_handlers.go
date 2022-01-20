package handlers

import (
	"go-gallery/src/core/ports"

	"github.com/gofiber/fiber/v2"
)

type FolderHandlers struct {
	folderService ports.FolderService
}

func NewFolderHandler(service ports.FolderService) *FolderHandlers {
	return &FolderHandlers{
		folderService: service,
	}
}

func (handler *FolderHandlers) Store(c *fiber.Ctx) error {
	title := c.Params("title")
	err := handler.folderService.Store(title)
	if err != nil {
		return err
	}
	return nil
}
