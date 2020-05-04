package focus

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	Create(user *User) (*User, error)
	FindUserByEmail(email string) (*User, error)
	ValidateEmailAndPassword(email, password string) (bool, error)
}
