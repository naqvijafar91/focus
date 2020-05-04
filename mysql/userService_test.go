package mysql

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/naqvijafar91/focus"
)

func createConnection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createUserService(t *testing.T) *UserService {
	db, err := createConnection()
	if err != nil {
		t.Error(err)
		return nil
	}
	usr, err := NewUserService(db)
	if err != nil {
		t.Error(err)
		return nil
	}
	return usr
}
func TestCreate(t *testing.T) {
	usr := createUserService(t)
	if usr == nil {
		return
	}
	_, err := usr.Create(&focus.User{Email: "xyz@xx.com", Password: "xxx"})
	if err != nil {
		t.Error("Should not throw error")
	}
}

func TestFindUserByEmail(t *testing.T) {
	usr := createUserService(t)
	if usr == nil {
		return
	}
	_, err := usr.Create(&focus.User{Email: "xyz@xx.com", Password: "xxx"})
	user, err := usr.FindUserByEmail("xyz@xx.com")
	if err != nil {
		t.Error(err)
		return
	}
	if user.Email != "xyz@xx.com" {
		t.Error("Found incorrect email")
	}
}

func TestNoDuplicateEmail(t *testing.T) {
	usr := createUserService(t)
	if usr == nil {
		return
	}
	_, err := usr.Create(&focus.User{Email: "xyz@xx.com", Password: "xxx"})
	if err != nil {
		t.Error("Should not throw error")
		return
	}
	_, err = usr.Create(&focus.User{Email: "xyz@xx.com", Password: "xxx"})
	if err == nil {
		t.Error("Should throw error if duplicate emails are inserted")
	}
}

func TestValidateEmailAndPassword(t *testing.T) {
	usr := createUserService(t)
	if usr == nil {
		return
	}
	_, err := usr.Create(&focus.User{Email: "xyz@xx.com", Password: "xxx"})
	if err != nil {
		t.Error("Should not throw error")
	}
	validated, err := usr.ValidateEmailAndPassword("xyz@xx.com", "xxx")
	if err != nil {
		t.Error("Should not throw error ", err)
		return
	}
	if !validated {
		t.Error("Failed to validate with correct credentials")
	}
}

func TestShouldNotValidateEmailAndPasswordPart1(t *testing.T) {
	usr := createUserService(t)
	if usr == nil {
		return
	}
	_, err := usr.Create(&focus.User{Email: "xyz@xx.com", Password: "xxx"})
	if err != nil {
		t.Error("Should not throw error")
	}
	validated, err := usr.ValidateEmailAndPassword("xyz@xx.com", "xxy")
	if err == nil {
		t.Error("Should throw error if password is different")
		return
	}
	if validated {
		t.Error("should not validate without correct credentials")
	}
}

func TestShouldNotValidateEmailAndPasswordPart2(t *testing.T) {
	usr := createUserService(t)
	if usr == nil {
		return
	}
	_, err := usr.Create(&focus.User{Email: "xyz@xx.com", Password: "xxx"})
	if err != nil {
		t.Error("Should not throw error")
	}
	validated, err := usr.ValidateEmailAndPassword("xyuy@xx.com", "xxx")
	if err == nil {
		t.Error("Should throw error if email is different")
		return
	}
	if validated {
		t.Error("should not validate without correct credentials")
	}
}