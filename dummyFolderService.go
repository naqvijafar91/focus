package main

import "github.com/google/uuid"

type DummyFolderService struct {
	folders []*Folder
}

func (dfs *DummyFolderService) Create(folder *Folder) (*Folder, error) {
	newFolder := &Folder{ID: uuid.New().String(), Name: folder.Name, UserID: folder.UserID}
	dfs.folders = append(dfs.folders, newFolder)
	return newFolder, nil
}

func (dfs *DummyFolderService) Update(folder *Folder) (*Folder, error) {
	return folder, nil
}

func (dfs *DummyFolderService) UpdateByID(ID string, folder *Folder) (*Folder, error) {
	return folder, nil
}
