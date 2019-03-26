package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/naqvijafar91/focus/handlers"
	"github.com/naqvijafar91/focus/memoryservices"
)

func main() {
	smux := http.NewServeMux()
	handlers.NewFolderHandler(&memoryservices.DummyFolderService{}).RegisterFolderRoutes(smux)
	handlers.NewHandler(&memoryservices.DummyUserService{}).RegisterUserRoutes(smux)
	handlers.NewTaskHandler(&memoryservices.DummyTaskService{}).RegisterTaskRoutes(smux)
	fmt.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", smux))
}
