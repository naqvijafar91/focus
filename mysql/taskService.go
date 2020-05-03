package mysql

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/naqvijafar91/focus"
)

type TaskService struct {
	db *gorm.DB
	fs focus.FolderService
}

func NewTaskService(db *gorm.DB, fs focus.FolderService) (*TaskService, error) {
	// Migrate the schema
	err := db.AutoMigrate(&focus.Task{}).Error
	if err != nil {
		return nil, err
	}
	return &TaskService{db, fs}, nil
}

func (ts *TaskService) Create(task *focus.Task) (*focus.Task, error) {
	// Validate if this folder exists or not
	folderID := task.FolderID
	folder, err := ts.fs.FindByID(folderID)
	if err != nil {
		return nil, err
	}
	if folder == nil {
		return nil, errors.New("Cannot create task without any valid folder")
	}
	newTask := &focus.Task{ID: uuid.New().String(), DueDate: task.DueDate,
		Description: task.Description, FolderID: task.FolderID}
	err = ts.db.Create(newTask).Error
	if err != nil {
		return nil, err
	}
	return newTask, nil
}

func (ts *TaskService) Update(task *focus.Task) (*focus.Task, error) {
	// Update single attribute if it is changed
	taskUpdated := &focus.Task{}
	err := ts.db.Where("id = ?", task.ID).First(taskUpdated).Error
	if err != nil {
		return nil, err
	}
	taskUpdated.Description = task.Description
	taskUpdated.DueDate = task.DueDate
	taskUpdated.CompletedDate = task.CompletedDate
	err = ts.db.Save(taskUpdated).Error
	if err != nil {
		return nil, err
	}
	return taskUpdated, nil
}

func (ts *TaskService) MarkAsComplete(taskID string) (*focus.Task, error) {
	taskUpdated := &focus.Task{}
	err := ts.db.Where("id = ?", taskID).First(taskUpdated).Error
	if err != nil {
		return nil, err
	}
	taskUpdated.CompletedDate = &focus.Time{time.Now()}
	err = ts.db.Save(taskUpdated).Error
	if err != nil {
		return nil, err
	}
	return taskUpdated, nil
}

func (ts *TaskService) GetAll() ([]*focus.Task, error) {
	var tasks []*focus.Task
	err := ts.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *TaskService) GetAllByFolderID(folderID string) ([]*focus.Task, error) {
	var tasks []*focus.Task
	err := ts.db.Where("folder_id = ?", folderID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
