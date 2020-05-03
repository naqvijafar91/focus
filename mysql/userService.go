package mysql

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/naqvijafar91/focus"
)

// UserService - Mysql backed user service
type UserService struct {
	db *gorm.DB
}

func NewMysqlConn(host string, port int, username string, dbName string, password string) (*gorm.DB, error) {
	// user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		username, password, host, port, dbName)
	db, err := gorm.Open("mysql", sqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewUserService(db *gorm.DB) (*UserService, error) {
	// Migrate the schema
	err := db.AutoMigrate(&focus.User{}).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&focus.User{}).AddUniqueIndex("idx_user_email", "email").Error
	if err != nil {
		return nil, err
	}
	return &UserService{db}, nil
}

func (us *UserService) Create(user *focus.User) (*focus.User, error) {
	user.ID = uuid.New().String()
	err := us.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) FindUserByEmail(email string) (*focus.User, error) {
	usr := &focus.User{}
	err := us.db.Where("email = ?", email).First(usr).Error
	if err != nil {
		return nil, err
	}
	return usr, nil
}
