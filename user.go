package main

type User struct {
	ID       string
	Email    string
	Password string
}

type UserService interface {
	Create(user *User) (*User, error)
	FindUserByEmail(email string) (*User, error)
}
