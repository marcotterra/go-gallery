package main

import (
	"go-gallery/src/core/services"
	"go-gallery/src/handlers"
	"go-gallery/src/repositories"
	"go-gallery/src/server"
)

func main() {
	mongoConn := "secret"

	//repositories
	folderRepository, _ := repositories.NewFolderRepository(mongoConn)

	//services
	folderService := services.NewFolderService(folderRepository)

	//handlers
	folderHandlers := handlers.NewFolderHandler(folderService)

	server := server.NewServer(folderHandlers)
	server.Initialize()
}
