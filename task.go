package focus

import (
	"time"
)

type Task struct {
	ID            string    `json:"id"`
	Description   string    `json:"description"`
	FolderID      string    `json:"folder_id"`
	DueDate       time.Time `json:"due_date"`
	CompletedDate time.Time `json:"completed_date"`
}

type TaskService interface {
	Create(task *Task) (*Task, error)
	GetAll() ([]*Task, error)
	GetAllByFolderID(folderID string) ([]*Task, error)
	MarkAsComplete(taskID string) (*Task, error)
	Update(task *Task) (*Task, error)
}
