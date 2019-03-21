package main

type Folder struct {
	ID     string
	Name   string
	UserID string
}

type FolderService interface {
	Create(folder *Folder) (*Folder, error)
	Update(folder *Folder) (*Folder, error)
	UpdateByID(ID string, folder *Folder) (*Folder, error)
}
