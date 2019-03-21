package main

type User struct {
	ID       int
	Email    string
	Password string
}

type UserService interface {
	Create(user *User) (*User, error)
}

type userService struct {
}

func (us *userService) Create(user *User) (*User, error) {
	return &User{ID: 22, Email: user.Email, Password: user.Password}, nil
}
