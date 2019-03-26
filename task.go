package focus

import "time"

type Task struct {
	ID            string
	Description   string
	DueDate       time.Time
	CompletedDate time.Time
}

type TaskService interface {
	Create(task *Task) (*Task, error)
	GetAll() (*Task, error)
}
