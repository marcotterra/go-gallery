package services

import "go-gallery/src/core/ports"

type FolderService struct {
	folderRepository ports.FolderRepository
}

func NewFolderService(repository ports.FolderRepository) *FolderService {
	return &FolderService{
		folderRepository: repository,
	}
}

func (service *FolderService) Store(title string) error {
	err := service.folderRepository.Store(title)
	if err != nil {
		return err
	}
	return nil
}
