package mysql

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/naqvijafar91/focus"
)

type FolderService struct {
	db *gorm.DB
}

func NewFolderService(db *gorm.DB) (*FolderService, error) {
	// Migrate the schema
	err := db.AutoMigrate(&focus.Folder{}).Error
	if err != nil {
		return nil, err
	}
	return &FolderService{db}, nil
}
func (fs *FolderService) Create(folder *focus.Folder) (*focus.Folder, error) {
	newFolder := &focus.Folder{ID: uuid.New().String(), Name: folder.Name, UserID: folder.UserID}
	err := fs.db.Create(newFolder).Error
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

func (fs *FolderService) Update(folder *focus.Folder) (*focus.Folder, error) {
	// Update single attribute if it is changed
	folderUpdated := &focus.Folder{}
	err := fs.db.Where("id = ?", folder.ID).First(folderUpdated).Error
	if err != nil {
		return nil, err
	}
	folderUpdated.Name = folder.Name
	folderUpdated.UserID = folder.UserID
	err = fs.db.Save(folderUpdated).Error
	if err != nil {
		return nil, err
	}
	return folderUpdated, nil
}

func (fs *FolderService) UpdateByID(ID string, folder *focus.Folder) (*focus.Folder, error) {
	folder.ID = ID
	return fs.Update(folder)
}

func (fs *FolderService) GetAll() ([]*focus.Folder, error) {
	var folders []*focus.Folder
	err := fs.db.Find(&folders).Error
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func (fs *FolderService) GetAllByUserID(userID string) ([]*focus.Folder, error) {
	var folders []*focus.Folder
	err := fs.db.Where("user_id = ?", userID).Find(&folders).Error
	if err != nil {
		return nil, err
	}
	return folders, nil
}
