package main

import (
	"github.com/google/uuid"
)

type DummyUserService struct {
	users []User
}

func (us *DummyUserService) Create(user *User) (*User, error) {
	return &User{ID: uuid.New().String(), Email: user.Email, Password: user.Password}, nil
}

func (us *DummyUserService) FindUserByEmail(email string) (*User, error) {
	return &User{ID: "22", Email: "dummy@email.com", Password: "Password"}, nil
}
