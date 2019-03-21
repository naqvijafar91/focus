package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	smux := http.NewServeMux()
	NewFolderHandler(&DummyFolderService{}).RegisterFolderRoutes(smux)
	NewHandler(&DummyUserService{}).registerUserRoutes(smux)
	fmt.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", smux))
}
