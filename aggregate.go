package focus

type AggregateResponse struct {
	data []Data
}

type Data struct {
	FolderID       string
	Name           string
	RemainingTasks int
	Tasks          []*Task
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

	var response *AggregateResponse
	//Step2 - For every folder fetch task and create a data struct
	for i := 0; i < len(folders); i++ {
		tasksForFolder, err := agtr.ts.GetAllByFolderID(folders[i].ID)
		if err != nil {
			return nil, err
		}
		response.data = append(response.data, Data{
			FolderID:       folders[i].ID,
			Name:           folders[i].Name,
			RemainingTasks: 11,
			Tasks:          tasksForFolder})
	}
	return response, nil
}
