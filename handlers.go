package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handlers struct {
	userService UserService
}

func (handlers *Handlers) sendJSONResponse(w http.ResponseWriter, entity interface{}) {
	dataBytes, err := json.MarshalIndent(entity, "", " ")
	if err != nil {
		fmt.Fprintf(w, "There was an error")
	}
	fmt.Fprintln(w, string(dataBytes))
}

func (handlers *Handlers) userRegistration(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user *User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	savedUser, err := handlers.userService.Create(user)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}

	handlers.sendJSONResponse(w, savedUser)
}

func NewHandler(userService UserService) *Handlers {
	return &Handlers{userService}
}

func InitHandlerAndServeHttpServer(userService UserService) {
	handler := NewHandler(userService)
	handler.serveHttpServer()
}
func (handler *Handlers) serveHttpServer() {
	// http.HandleFunc("/", handler)
	fmt.Println("Server starting")
	http.HandleFunc("/user/register", handler.userRegistration)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
