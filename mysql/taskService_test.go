package mysql

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/naqvijafar91/focus"
)

func createTaskServiceWithConn(t *testing.T, conn *gorm.DB) *TaskService {
	taskServ, err := NewTaskService(conn, createFolderServiceWithConn(t, conn))
	if err != nil {
		t.Error(err)
		return nil
	}
	return taskServ
}
func createFolderServiceWithConn(t *testing.T, conn *gorm.DB) *FolderService {
	usr, err := NewFolderService(conn)
	if err != nil {
		t.Error(err)
		return nil
	}
	return usr
}

func createUserServiceWithConn(t *testing.T, conn *gorm.DB) *UserService {
	usr, err := NewUserService(conn)
	if err != nil {
		t.Error(err)
		return nil
	}
	return usr
}

func createUserWithConn(t *testing.T, email string, conn *gorm.DB) *focus.User {
	usr := createUserServiceWithConn(t, conn)
	if usr == nil {
		return nil
	}
	user, err := usr.Create(&focus.User{Email: email, Password: "xxx"})
	if err != nil {
		t.Error("Should not throw error")
		return nil
	}
	return user
}
func TestTaskCreate(t *testing.T) {
	conn, err := createConnection()
	if err != nil {
		t.Error(err)
		return
	}
	ts, fs, user := createTaskServiceWithConn(t, conn), createFolderServiceWithConn(t, conn), createUserWithConn(t, "dummy@xyz.com", conn)
	if ts == nil || fs == nil || user == nil {
		return
	}
	folder, err := createFolder("dummy folder", user, fs)
	if err != nil {
		t.Error(err)
		return
	}
	tsk, err := ts.Create(&focus.Task{Description: "Dummy Task", FolderID: folder.ID})
	if err != nil {
		t.Error(err)
		return
	}
	if tsk.Description != "Dummy Task" {
		t.Error("Task not created properly")
	}
}

func TestTaskShouldNotCreateWithoutFolderID(t *testing.T) {
	conn, err := createConnection()
	if err != nil {
		t.Error(err)
		return
	}
	ts, fs, user := createTaskServiceWithConn(t, conn), createFolderServiceWithConn(t, conn), createUserWithConn(t, "dummy@xyz.com", conn)
	if ts == nil || fs == nil || user == nil {
		return
	}
	tsk, err := ts.Create(&focus.Task{Description: "Dummy Task", FolderID: ""})
	if err == nil {
		t.Error("Should throw error if task is created without folder id")
		return
	}
	if tsk != nil {
		t.Error("Task should not be created")
	}
}

func TestTaskShouldNotCreateWithWrongFolderID(t *testing.T) {
	conn, err := createConnection()
	if err != nil {
		t.Error(err)
		return
	}
	ts, fs, user := createTaskServiceWithConn(t, conn), createFolderServiceWithConn(t, conn), createUserWithConn(t, "dummy@xyz.com", conn)
	if ts == nil || fs == nil || user == nil {
		return
	}
	tsk, err := ts.Create(&focus.Task{Description: "Dummy Task", FolderID: "695f1346-6edc-4eea-9fb5-4616a916b5d8"})
	if err == nil {
		t.Error("Should throw error if task is created wit wrong folder id")
		return
	}
	if tsk != nil {
		t.Error("Task should not be created")
	}
}

func TestUpdateTask(t *testing.T) {
	conn, err := createConnection()
	if err != nil {
		t.Error(err)
		return
	}
	ts, fs, user := createTaskServiceWithConn(t, conn), createFolderServiceWithConn(t, conn), createUserWithConn(t, "dummy@xyz.com", conn)
	if ts == nil || fs == nil || user == nil {
		return
	}
	folder, _ := fs.Create(&focus.Folder{Name: "ola", UserID: user.ID})
	tsk, _ := ts.Create(&focus.Task{Description: "Dummy Task", FolderID: folder.ID})
	tsk.Description = "Updated Task"
	updated, err := ts.Update(tsk)
	if err != nil {
		t.Error(err)
		return
	}
	if updated == nil {
		t.Error("Returned nil")
		return
	}
	if updated.Description != "Updated Task" {
		t.Error("Task not updated")
	}
}

func TestTaskGetAll(t *testing.T) {
	conn, err := createConnection()
	if err != nil {
		t.Error(err)
		return
	}
	ts, fs, user := createTaskServiceWithConn(t, conn), createFolderServiceWithConn(t, conn), createUserWithConn(t, "dummy@xyz.com", conn)
	if ts == nil || fs == nil || user == nil {
		return
	}
	folder, _ := fs.Create(&focus.Folder{Name: "ola", UserID: user.ID})
	for i := 0; i < 10; i++ {
		_, err := ts.Create(&focus.Task{Description: "Dummy Task", FolderID: folder.ID})
		if err != nil {
			t.Error("Should not throw error")
			return
		}
	}
	tasks, err := ts.GetAll()
	if err != nil {
		t.Error(err)
		return
	}
	if len(tasks) != 10 {
		t.Error("Less tasks found")
	}
}

func TestTaskGetAllByFolderID(t *testing.T) {
	conn, err := createConnection()
	if err != nil {
		t.Error(err)
		return
	}
	ts, fs, user := createTaskServiceWithConn(t, conn), createFolderServiceWithConn(t, conn), createUserWithConn(t, "dummy@xyz.com", conn)
	if ts == nil || fs == nil || user == nil {
		return
	}
	folder, _ := fs.Create(&focus.Folder{Name: "ola", UserID: user.ID})
	folder2, _ := fs.Create(&focus.Folder{Name: "ola 2", UserID: user.ID})
	for i := 0; i < 10; i++ {
		desc := fmt.Sprintf("desc-%d", i)
		_, err := ts.Create(&focus.Task{Description: desc, FolderID: folder.ID})
		_, err = ts.Create(&focus.Task{Description: desc, FolderID: folder2.ID})
		if err != nil {
			t.Error("Should not throw error")
			return
		}
	}
	tasks1, err1 := ts.GetAllByFolderID(folder.ID)
	tasks2, err2 := ts.GetAllByFolderID(folder2.ID)
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
		return
	}
	if len(tasks1) != 10 || len(tasks2) != 10 {
		t.Error("Less Tasks found")
		return
	}
	for i := 0; i < 10; i++ {
		desc := fmt.Sprintf("desc-%d", i)
		if tasks1[i].Description != desc || tasks2[i].Description != desc {
			t.Error("Name not matching at index", i)
		}
	}
}
