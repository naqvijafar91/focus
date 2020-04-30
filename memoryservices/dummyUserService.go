package memoryservices

import (
	"errors"

	"github.com/google/uuid"
	"github.com/naqvijafar91/focus"
)

type DummyUserService struct {
	users []focus.User
}

func (us *DummyUserService) Create(user *focus.User) (*focus.User, error) {
	if _, err := us.FindUserByEmail(user.Email); err == nil {
		return nil, errors.New("User already exists")
	}
	usr := &focus.User{ID: uuid.New().String(), Email: user.Email, Password: user.Password}
	us.users = append(us.users, *usr)
	return usr, nil
}

func (us *DummyUserService) FindUserByEmail(email string) (*focus.User, error) {
	for _, usr := range us.users {
		if usr.Email == email {
			return &usr, nil
		}
	}
	return nil, errors.New("No user Found")
}
