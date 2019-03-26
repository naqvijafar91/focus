package memoryservices

import (
	"github.com/google/uuid"
	"github.com/naqvijafar91/focus"
)

type DummyUserService struct {
	users []focus.User
}

func (us *DummyUserService) Create(user *focus.User) (*focus.User, error) {
	return &focus.User{ID: uuid.New().String(), Email: user.Email, Password: user.Password}, nil
}

func (us *DummyUserService) FindUserByEmail(email string) (*focus.User, error) {
	return &focus.User{ID: "22", Email: "dummy@email.com", Password: "Password"}, nil
}
