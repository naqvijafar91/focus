package main

type User struct {
	ID       int
	Email    string
	Password string
}

type UserService interface {
	Create(user *User) (*User, error)
	FindUserByEmail(email string) (*User, error)
}
