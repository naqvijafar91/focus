package memoryservices

import (
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
	for _, folderInStore := range dfs.folders {
		if folder.ID == folderInStore.ID {
			folderInStore = folder
		}
	}
	return folder, nil
}

func (dfs *DummyFolderService) UpdateByID(ID string, folder *focus.Folder) (*focus.Folder, error) {
	folder.ID = ID
	return dfs.Update(folder)
}

func (dfs *DummyFolderService) GetAll() ([]*focus.Folder, error) {
	return dfs.folders, nil
}
