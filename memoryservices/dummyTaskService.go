package memoryservices

import (
	"time"

	"github.com/google/uuid"
	"github.com/naqvijafar91/focus"
)

type DummyTaskService struct {
	tasks []*focus.Task
}

func (dts *DummyTaskService) Create(task *focus.Task) (*focus.Task, error) {
	newTask := &focus.Task{ID: uuid.New().String(), DueDate: task.DueDate, Description: task.Description}
	return newTask, nil
}

func (dts *DummyTaskService) Update(task *focus.Task) (*focus.Task, error) {
	for i := 0; i < len(dts.tasks); i++ {
		if task.ID == dts.tasks[i].ID {
			dts.tasks[i] = task
			return dts.tasks[i], nil
		}
	}
	return task, nil
}

func (dts *DummyTaskService) MarkAsComplete(taskID string) (*focus.Task, error) {
	var foundTask *focus.Task
	for i := 0; i < len(dts.tasks); i++ {
		if taskID == dts.tasks[i].ID {
			dts.tasks[i].CompletedDate = time.Now()
			foundTask = dts.tasks[i]
			return foundTask, nil
		}
	}
	return foundTask, nil
}

func (dts *DummyTaskService) GetAll() ([]*focus.Task, error) {
	return dts.tasks, nil
}