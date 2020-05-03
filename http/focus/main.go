package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/naqvijafar91/focus"
	"github.com/rs/cors"

	"github.com/naqvijafar91/focus/handlers"
	"github.com/naqvijafar91/focus/mysql"
)

func main() {
	smux := http.NewServeMux()
	// These are memory backed services
	// folderService := &memorybackedservices.DummyFolderService{}
	// userService := &memorybackedservices.DummyUserService{}
	// taskService := &memorybackedservices.DummyTaskService{}
	userService, folderService, taskService := initServices()
	handlers.NewFolderHandler(folderService).RegisterFolderRoutes(smux)
	handlers.NewHandler(userService).RegisterUserRoutes(smux)
	handlers.NewTaskHandler(taskService).RegisterTaskRoutes(smux)
	aggregatorService := focus.NewAggregatorService(taskService,
		folderService,
		userService)
	handlers.NewAggregatorHandler(aggregatorService).RegisterAggregatorRoutes(smux)
	fmt.Println("Server starting")
	handler := cors.AllowAll().Handler(smux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func initServices() (focus.UserService, focus.FolderService, focus.TaskService) {
	db, err := mysql.NewMysqlConn("localhost", 3306, "focus", "focus", "pwd")
	if err != nil {
		panic(err)
	}
	userService, err := mysql.NewUserService(db)
	folderService, err := mysql.NewFolderService(db)
	taskService, err := mysql.NewTaskService(db, folderService)
	if err != nil {
		panic(err)
	}
	return userService, folderService, taskService
}
