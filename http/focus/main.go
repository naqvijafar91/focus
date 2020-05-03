package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/naqvijafar91/focus"
	"github.com/rs/cors"

	"github.com/naqvijafar91/focus/handlers"
	"github.com/naqvijafar91/focus/memorybackedservices"
)

func main() {
	smux := http.NewServeMux()
	folderService := &memorybackedservices.DummyFolderService{}
	userService := &memorybackedservices.DummyUserService{}
	taskService := &memorybackedservices.DummyTaskService{}
	handlers.NewFolderHandler(folderService).RegisterFolderRoutes(smux)
	handlers.NewHandler(userService).RegisterUserRoutes(smux)
	handlers.NewTaskHandler(taskService).RegisterTaskRoutes(smux)
	aggregatorService := focus.NewAggregatorService(taskService,
		folderService,
		userService)
	handlers.NewAggregatorHandler(aggregatorService).RegisterAggregatorRoutes(smux)
	fmt.Println("Server starting")
	// handler := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedHeaders: []string{"Authorization"}}).Handler(smux)
	handler := cors.AllowAll().Handler(smux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
