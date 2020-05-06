package mysql

import (
	"errors"
	"strings"

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
	// First check if a inbox folder exists or not, if exists, then append a number at the end
	verifiedFolder, err := fs.updateFolderNameIfItIsInbox(folder)
	if err != nil {
		return nil, err
	}
	newFolder := &focus.Folder{ID: uuid.New().String(), Name: verifiedFolder.Name, UserID: folder.UserID}
	err = fs.db.Create(newFolder).Error
	if err != nil {
		return nil, err
	}
	return newFolder, nil
}

func (fs *FolderService) updateFolderNameIfItIsInbox(folder *focus.Folder) (*focus.Folder, error) {
	if strings.Trim(folder.Name, " ") != "Inbox" {
		return folder, nil
	}
	if strings.Trim(folder.UserID, " ") == "" {
		return nil, errors.New("User ID of folder cannot be empty")
	}
	inboxFolderExists, err := fs.FindInboxFolder(folder.UserID)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if inboxFolderExists != nil && inboxFolderExists.Name == "Inbox" {
		folder.Name = folder.Name + " 2"
	}
	return folder, nil
}

func (fs *FolderService) Update(folder *focus.Folder) (*focus.Folder, error) {
	verifiedFolder, err := fs.updateFolderNameIfItIsInbox(folder)
	if err != nil {
		return nil, err
	}
	// Update single attribute if it is changed
	folderUpdated := &focus.Folder{}
	err = fs.db.Where("id = ?", folder.ID).First(folderUpdated).Error
	if err != nil {
		return nil, err
	}
	// Inbox folder cannot be updated
	if folderUpdated.Name == "Inbox" {
		return nil, errors.New("Cannot update Inbox folder")
	}
	folderUpdated.Name = verifiedFolder.Name
	err = fs.db.Save(folderUpdated).Error
	if err != nil {
		return nil, err
	}
	return folderUpdated, nil
}

func (fs *FolderService) UpdateByID(ID string, folder *focus.Folder) (*focus.Folder, error) {
	foundFolder, err := fs.FindByID(ID)
	if err != nil {
		return nil, err
	}
	// Update the folder's user id with the found folder
	if foundFolder == nil {
		return nil, errors.New("record not found")
	}
	folder.UserID = foundFolder.UserID
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

func (fs *FolderService) FindInboxFolder(userID string) (*focus.Folder, error) {
	folder := &focus.Folder{}
	err := fs.db.Where("name = ? and user_id=?", "Inbox", userID).First(folder).Error
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (fs *FolderService) FindByID(folderID string) (*focus.Folder, error) {
	folder := &focus.Folder{}
	err := fs.db.Where("id = ?", folderID).First(folder).Error
	if err != nil {
		return nil, err
	}
	return folder, nil
}
