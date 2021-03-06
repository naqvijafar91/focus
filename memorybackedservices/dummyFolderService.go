package memorybackedservices

import (
	"errors"

	"github.com/google/uuid"
	"github.com/naqvijafar91/focus"
)

type DummyFolderService struct {
	folders []*focus.Folder
}

func (dfs *DummyFolderService) Create(folder *focus.Folder) (*focus.Folder, error) {
	newFolder := &focus.Folder{ID: uuid.New().String(), Name: folder.Name, UserID: folder.UserID}
	dfs.folders = append(dfs.folders, newFolder)
	return newFolder, nil
}

func (dfs *DummyFolderService) Update(folder *focus.Folder) (*focus.Folder, error) {
	for i := 0; i < len(dfs.folders); i++ {
		folderInStore := dfs.folders[i]
		if folder.ID == folderInStore.ID {
			dfs.folders[i] = folder
			return dfs.folders[i], nil
		}
	}
	return nil, errors.New("Folder not found")
}

func (dfs *DummyFolderService) UpdateByID(ID string, folder *focus.Folder) (*focus.Folder, error) {
	folder.ID = ID
	return dfs.Update(folder)
}

func (dfs *DummyFolderService) GetAll() ([]*focus.Folder, error) {
	return dfs.folders, nil
}

func (dfs *DummyFolderService) GetAllByUserID(userID string) ([]*focus.Folder, error) {
	return dfs.folders, nil
}

func (dfs *DummyFolderService) FindByID(folderID string) (*focus.Folder, error) {
	return nil, nil
}
