package memorybackedservices

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/naqvijafar91/focus"
)

type DummyTaskService struct {
	tasks []*focus.Task
}

func (dts *DummyTaskService) Create(task *focus.Task) (*focus.Task, error) {
	newTask := &focus.Task{ID: uuid.New().String(), DueDate: task.DueDate,
		Description: task.Description, FolderID: task.FolderID}
	dts.tasks = append(dts.tasks, newTask)
	return newTask, nil
}

func (dts *DummyTaskService) Update(task *focus.Task) (*focus.Task, error) {
	for i := 0; i < len(dts.tasks); i++ {
		if task.ID == dts.tasks[i].ID {
			dts.tasks[i] = task
			return dts.tasks[i], nil
		}
	}
	return nil, errors.New("Task Not Found")
}

func (dts *DummyTaskService) MarkAsComplete(taskID string) (*focus.Task, error) {
	var foundTask *focus.Task
	for i := 0; i < len(dts.tasks); i++ {
		if taskID == dts.tasks[i].ID {
			dts.tasks[i].CompletedDate = &focus.Time{time.Now()}
			foundTask = dts.tasks[i]
			return foundTask, nil
		}
	}
	return foundTask, nil
}

func (dts *DummyTaskService) GetAll() ([]*focus.Task, error) {
	return dts.tasks, nil
}

func (dts *DummyTaskService) GetAllByFolderID(folderID string) ([]*focus.Task, error) {
	return dts.tasks, nil
}
