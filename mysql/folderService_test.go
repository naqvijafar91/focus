package mysql

import (
	"fmt"
	"testing"

	"github.com/naqvijafar91/focus"
)

func createFolderService(t *testing.T) *FolderService {
	db, err := createConnection()
	if err != nil {
		t.Error(err)
		return nil
	}
	usr, err := NewFolderService(db)
	if err != nil {
		t.Error(err)
		return nil
	}
	return usr
}

func createUser(t *testing.T, email string) *focus.User {
	usr := createUserService(t)
	if usr == nil {
		return nil
	}
	user, err := usr.Create(&focus.User{Email: email, LoginCode: "xxx"})
	if err != nil {
		t.Error("Should not throw error")
		return nil
	}
	return user
}

func createFolder(name string, user *focus.User, fs *FolderService) (*focus.Folder, error) {
	return fs.Create(&focus.Folder{Name: name, UserID: user.ID})
}

func TestFolderCreate(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	folder, err := fs.Create(&focus.Folder{Name: "ola", UserID: user.ID})
	if err != nil {
		t.Error("Folder is nil")
		return
	}
	if folder.Name != "ola" && folder.UserID != user.ID {
		t.Error("Wrong user ID and name")
	}
}

func TestFolderCreateInboxDuplicate(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	_, err := fs.Create(&focus.Folder{Name: "Inbox", UserID: user.ID})
	if err != nil {
		t.Error("Folder is nil")
		return
	}
	folder, err := fs.Create(&focus.Folder{Name: "Inbox", UserID: user.ID})
	if err != nil {
		t.Error("Folder is nil")
		return
	}
	if folder.Name != "Inbox 2" || folder.UserID != user.ID {
		t.Error("Wrong user ID or name")
	}
}

func TestUpdateFolder(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	folder, _ := fs.Create(&focus.Folder{Name: "ola", UserID: user.ID})
	folder.Name = "updated Name"
	updated, err := fs.Update(folder)
	if err != nil {
		t.Error(err)
		return
	}
	if updated.Name != "updated Name" {
		t.Error("Name not updated")
	}
}

func TestShouldNotUpdateWithNameInbox(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	_, err := fs.Create(&focus.Folder{Name: "Inbox", UserID: user.ID})
	if err != nil {
		t.Error("Folder is nil")
		return
	}
	folder, _ := fs.Create(&focus.Folder{Name: "ola", UserID: user.ID})
	folder.Name = "Inbox"
	updated, err := fs.Update(folder)
	if err != nil {
		t.Error(err)
		return
	}
	if updated.Name != "Inbox 2" {
		t.Error("Name not updated to Inbox 2")
	}
}

func TestShouldNotUpdateWithInboxWithAnyOtherName(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	folder, err := fs.Create(&focus.Folder{Name: "Inbox", UserID: user.ID})
	if err != nil {
		t.Error("Folder is nil")
		return
	}
	folder.Name = "Any random name"
	updated, err := fs.Update(folder)
	if err == nil {
		t.Error("Should throw error")
		return
	}
	if updated != nil {
		t.Error("Name should not be updated")
	}
}

func TestUpdateFolderShouldNotUpdateUserID(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	folder, _ := fs.Create(&focus.Folder{Name: "ola", UserID: user.ID})
	folder.Name = "updated Name"
	folder.UserID = "Dummy ID"
	updated, err := fs.Update(folder)
	if err != nil {
		t.Error(err)
		return
	}
	if updated.Name != "updated Name" {
		t.Error("Name not updated")
	}
	if updated.UserID != user.ID {
		t.Error("Should not update user id of folder")
	}
}

func TestGetAll(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	// Create 10 folders
	for i := 0; i < 10; i++ {
		_, err := createFolder(fmt.Sprintf("name-%d", i), user, fs)
		if err != nil {
			t.Error(err)
			return
		}
	}
	folders, err := fs.GetAll()
	if err != nil {
		t.Error(err)
		return
	}
	if len(folders) != 10 {
		t.Error("Less folders found")
	}
}

func TestGetByID(t *testing.T) {
	fs, user := createFolderService(t), createUser(t, "xyz@ss.com")
	if fs == nil || user == nil {
		return
	}
	folder, err := createFolder(fmt.Sprintf("name-%d", 786), user, fs)
	fetched, err := fs.FindByID(folder.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if fetched.Name != fmt.Sprintf("name-%d", 786) {
		t.Error("Failed to fetch by id")
	}
}

func TestGetAllByUserID(t *testing.T) {
	fs, user1, user2 := createFolderService(t), createUser(t, "xyz@ss.com"), createUser(t, "xya@ss.com")
	if fs == nil || user1 == nil || user2 == nil {
		return
	}
	// Create 10 folders for user 1 and user 2 respectively
	for i := 0; i < 10; i++ {
		_, err := createFolder(fmt.Sprintf("name-%d", i), user1, fs)
		_, err = createFolder(fmt.Sprintf("name-%d", i), user2, fs)
		if err != nil {
			t.Error(err)
			return
		}
	}
	folders1, err1 := fs.GetAllByUserID(user1.ID)
	folders2, err2 := fs.GetAllByUserID(user2.ID)
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
		return
	}
	if len(folders1) != 10 || len(folders2) != 10 {
		t.Error("Less folders found")
		return
	}
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("name-%d", i)
		if folders1[i].Name != name || folders2[i].Name != name {
			t.Error("Name not matching at index", i)
		}
	}
}

func TestGetAllByUserIDShouldNotThrowErrorIfEmpty(t *testing.T) {
	fs, user1, user2 := createFolderService(t), createUser(t, "xyz@ss.com"), createUser(t, "xya@ss.com")
	if fs == nil || user1 == nil || user2 == nil {
		return
	}
	folders1, err1 := fs.GetAllByUserID(user1.ID)
	folders2, err2 := fs.GetAllByUserID(user2.ID)
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
		return
	}
	if len(folders1) != 0 || len(folders2) != 0 {
		t.Error("Length should be 0")
		return
	}
}
