package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/naqvijafar91/focus"
)

type Handlers struct {
	userService focus.UserService
}

//@Todo:Find a way to properly handle errors
func (handlers *Handlers) sendError(w http.ResponseWriter, err error) {

}

func (handlers *Handlers) userLogin(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userLoginReq *focus.User
	err := decoder.Decode(&userLoginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	valid, err := handlers.userService.ValidateEmailAndPassword(userLoginReq.Email, userLoginReq.Password)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	if !valid {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Its an error %s", errors.New("Invalid Password"))
		return
	}
	// Since it is a valid user, find and query
	foundUser, err := handlers.userService.FindUserByEmail(userLoginReq.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       foundUser.ID,
		"username": foundUser.Email,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, error)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
		"user":  valid})
}

func (handlers *Handlers) userRegistration(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user *focus.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	savedUser, err := handlers.userService.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       savedUser.ID,
		"username": savedUser.Email,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
		"user":  savedUser})
}

func NewHandler(userService focus.UserService) *Handlers {
	return &Handlers{userService}
}

func (handlers *Handlers) RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user/register", handlers.userRegistration)
	mux.HandleFunc("/user/login", handlers.userLogin)
}
