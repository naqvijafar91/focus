package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/naqvijafar91/focus"
)

// Handlers - Login/Registration would follow the same flow, and it will be a 2 step process everytime
type Handlers struct {
	userLoginService focus.UserLoginService
}

//@Todo:Find a way to properly handle errors
func (handlers *Handlers) sendError(w http.ResponseWriter, err error) {

}

// verify will check if the code is correct and generate a Bearer token in return
func (handlers *Handlers) verify(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userLoginReq *focus.User
	err := decoder.Decode(&userLoginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	valid, err := handlers.userLoginService.ValidateLoginCodeAndInit(userLoginReq.Email, userLoginReq.LoginCode)
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
	foundUser, err := handlers.userLoginService.FindUserByEmail(userLoginReq.Email)
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

// generate will generate a code and create the user if it not exists and share the code with the user over email
func (handlers *Handlers) generate(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user *focus.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	_, err = handlers.userLoginService.GenerateAndShareCode(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Its an error %s", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Please check the email for code"})
}

func NewHandler(userLoginService focus.UserLoginService) *Handlers {
	return &Handlers{userLoginService}
}

func (handlers *Handlers) RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user/generate", handlers.generate)
	mux.HandleFunc("/user/verify", handlers.verify)
}
