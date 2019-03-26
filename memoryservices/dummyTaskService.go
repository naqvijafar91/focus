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
	for _, taskInStore := range dts.tasks {
		if taskInStore.ID == task.ID {
			taskInStore = task
		}
	}
	return task, nil
}

func (dts *DummyTaskService) MarkAsComplete(taskID string) (*focus.Task, error) {
	var completedTask *focus.Task
	for _, taskInStore := range dts.tasks {
		if taskInStore.ID == taskID {
			taskInStore.CompletedDate = time.Now()
		}
	}
	return completedTask, nil
}

func (dts *DummyTaskService) GetAll() ([]*focus.Task, error) {
	return dts.tasks, nil
}
