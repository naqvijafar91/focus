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
	tasks, err := th.taskService.GetAll()
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tasks": tasks})
}

func (th *TaskHandler) MarkCompleted(w http.ResponseWriter, req *http.Request) {

}

func NewTaskHandler(ts focus.TaskService) *TaskHandler {
	return &TaskHandler{ts}
}

func (th *TaskHandler) handleTaskRoutes(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		th.GetAll(w, req)
		break
	case http.MethodPost:
		th.Create(w, req)
		break
	}
}

func (th *TaskHandler) RegisterFolderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/task", th.handleTaskRoutes)
	mux.HandleFunc("/task/complete", th.MarkCompleted)
}
