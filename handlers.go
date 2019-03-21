package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type userLoginRequest struct {
	Email    string
	Password string
}

type Handlers struct {
	userService UserService
}

func (handlers *Handlers) sendJSON(w http.ResponseWriter, entity interface{}) {
	dataBytes, err := json.MarshalIndent(entity, "", " ")
	if err != nil {
		fmt.Fprintf(w, "There was an error")
	}
	fmt.Fprintln(w, string(dataBytes))
}

//@Todo:Find a way to properly handle errors
func (handlers *Handlers) sendError(w http.ResponseWriter, err error) {

}

func (handlers *Handlers) userLogin(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userLoginReq *userLoginRequest
	err := decoder.Decode(&userLoginReq)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	foundUser, err := handlers.userService.FindUserByEmail(userLoginReq.Email)
	if err != nil {
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}

	if foundUser.Password != userLoginReq.Password {
		fmt.Fprintf(w, "Its an error %s", errors.New("Invalid password"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": foundUser.Email,
		"password": foundUser.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
		"user":  foundUser})
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

	handlers.sendJSON(w, savedUser)
}

func NewHandler(userService UserService) *Handlers {
	return &Handlers{userService}
}

func InitHandlerAndServeHttpServer(userService UserService) {
	handler := NewHandler(userService)
	handler.serveHTTPServer()
}
func (handlers *Handlers) serveHTTPServer() {
	// http.HandleFunc("/", handler)
	fmt.Println("Server starting")
	http.HandleFunc("/user/register", handlers.userRegistration)
	http.HandleFunc("/user/login", handlers.userLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
