package focus

import (
	"fmt"
	"time"
)

type AggregateResponse struct {
	Data []Data `json:"data"`
}

type Data struct {
	FolderID       string  `json:"id"`
	Name           string  `json:"name"`
	RemainingTasks int     `json:"remaining_tasks"`
	Tasks          []*Task `json:"tasks"`
}

type AggregatorService interface {
	GetAllData(userID string) (*AggregateResponse, error)
}

type Aggregator struct {
	ts TaskService
	fs FolderService
	us UserService
}

// GetAllData - Fetch AggregateResponse for a user
func (agtr *Aggregator) GetAllData(userID string) (*AggregateResponse, error) {

	//Step 1 - Fetch All Folders for the user
	folders, err := agtr.fs.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := &AggregateResponse{}
	//Step2 - For every folder fetch task and create a data struct
	for i := 0; i < len(folders); i++ {
		fmt.Println("Filling Folder with name", folders[i].Name)
		tasksForFolder, err := agtr.ts.GetAllByFolderID(folders[i].ID)
		if err != nil {
			return nil, err
		}

		if tasksForFolder == nil {
			tasksForFolder = make([]*Task, 0)
		}
		// Find out the remaining tasks in the folder
		// Remaining tasks = Tasks whose completed date is after the current date or nil
		remainingTasks := 0
		for _, task := range tasksForFolder {
			if task.CompletedDate == nil || task.CompletedDate.After(time.Now()) {
				remainingTasks++
			}
		}
		response.Data = append(response.Data, Data{
			FolderID:       folders[i].ID,
			Name:           folders[i].Name,
			RemainingTasks: remainingTasks,
			Tasks:          tasksForFolder})
	}
	return response, nil
}

func NewAggregatorService(ts TaskService, fs FolderService, us UserService) *Aggregator {
	return &Aggregator{ts, fs, us}
}
