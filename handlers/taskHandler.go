package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naqvijafar91/focus"
)

type TaskHandler struct {
	taskService focus.TaskService
}

func (th *TaskHandler) Create(w http.ResponseWriter, req *http.Request) {
	var newTask *focus.Task
	err := json.NewDecoder(req.Body).Decode(&newTask)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse request %s", err)
		return
	}
	createdTask, err := th.taskService.Create(newTask)
	if err != nil {
		fmt.Fprintf(w, "Failed to save task %s", err)
		return
	}
	json.NewEncoder(w).Encode(createdTask)
}

func (th *TaskHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	tasks, err := th.taskService.GetAllByFolderID("dummyId")
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tasks": tasks})
}

func (th *TaskHandler) MarkCompleted(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	taskID := req.Form["taskID"][0]
	completedTask, err := th.taskService.MarkAsComplete(taskID)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(completedTask)
}

func (th *TaskHandler) Update(w http.ResponseWriter, req *http.Request) {
	var updatedTaskInReq *focus.Task
	err := json.NewDecoder(req.Body).Decode(&updatedTaskInReq)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	updatedTaskFromDB, err := th.taskService.Update(updatedTaskInReq)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(updatedTaskFromDB)
}

func NewTaskHandler(ts focus.TaskService) *TaskHandler {
	return &TaskHandler{ts}
}

func (th *TaskHandler) RegisterTaskRoutes(mux *http.ServeMux) {
	middlewares := chainMiddleware(withUserParsing)
	mux.HandleFunc("/task", middlewares(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			//@Note: This method has to be removed
			th.GetAll(w, req)
			break
		case http.MethodPost:
			th.Create(w, req)
			break
		case http.MethodPut:
			th.Update(w, req)
			break
		}
	}))
	mux.HandleFunc("/task/complete", th.MarkCompleted)
}
