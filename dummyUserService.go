package main

type DummyUserService struct {
}

func (us *DummyUserService) Create(user *User) (*User, error) {
	return &User{ID: 22, Email: user.Email, Password: user.Password}, nil
}

func (us *DummyUserService) FindUserByEmail(email string) (*User, error) {
	return &User{ID: 22, Email:"dummy@email.com", Password: "Password"}, nil
}
