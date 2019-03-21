package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
		fmt.Fprintln(w, error)
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": savedUser.Email,
		"password": savedUser.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
		"user":  savedUser})
}

func NewHandler(userService UserService) *Handlers {
	return &Handlers{userService}
}

func (handlers *Handlers) registerUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user/register", handlers.userRegistration)
	mux.HandleFunc("/user/login", handlers.userLogin)
}
