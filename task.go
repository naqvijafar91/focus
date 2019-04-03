package focus

import (
	"time"
)

type Task struct {
	ID            string
	Description   string
	FolderID      string
	DueDate       time.Time
	CompletedDate time.Time
}

type TaskService interface {
	Create(task *Task) (*Task, error)
	GetAll() ([]*Task, error)
	GetAllByFolderID(folderID string) ([]*Task, error)
	MarkAsComplete(taskID string) (*Task, error)
	Update(task *Task) (*Task, error)
}
