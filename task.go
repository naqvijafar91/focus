package focus

import "time"

type Task struct {
	ID int
	dueDate time.Time
	completedDate time.Time
}

type TaskService interface {
	

}