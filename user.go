package focus

import (
	"math/rand"
	"strconv"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	LoginCode string `json:"login_code"`
}

// LoginCodeNotificationService is responsible for sharing the login code generated with the user
type LoginCodeNotificationService interface {
	SendCodeOnEmail(email, code string) error
}

type CodeGenerator interface {
	Generate() string
}

type UserService interface {
	Create(user *User) (*User, error)
	FindUserByEmail(email string) (*User, error)
	UpdateLoginCode(email, code string) error
	ValidateEmailAndLoginCode(email, loginCode string) (bool, error)
}

// UserLoginService is responsible for handling the login operation for the app, Login/Registration is a
// 2 step process, first step will generate the code and share it with user, next step is to validate the
// code
type UserLoginService interface {
	GenerateAndShareCode(email string) (string, error)
	ValidateLoginCodeAndInit(email, code string) (bool, error)
	FindUserByEmail(email string) (*User, error)
}

type userLoginService struct {
	notificationService LoginCodeNotificationService
	userService         UserService
	codeGenerator       CodeGenerator
	folderService       FolderService
}

// NewUserLoginService constructor for UserLoginService
func NewUserLoginService(notificationService LoginCodeNotificationService, userService UserService,
	codeGenerator CodeGenerator, folderService FolderService) UserLoginService {
	return &userLoginService{notificationService, userService, codeGenerator, folderService}
}

// GenerateAndShareCode will create a user by this email if not exists and also updates the login code
// then it shares the code with the user via notification service
func (uls *userLoginService) GenerateAndShareCode(email string) (string, error) {
	// Check if this user exists or not
	usr, _ := uls.FindUserByEmail(email)
	// Generate 4 digit numeric code
	code := uls.codeGenerator.Generate()
	if usr == nil {
		// Create this user
		createdUser, err := uls.userService.Create(&User{Email: email, LoginCode: code})
		if err != nil {
			return "", err
		}
		usr = createdUser
	} else {
		// Just update the login code
		err := uls.userService.UpdateLoginCode(email, code)
		if err != nil {
			return "", err
		}
	}
	// Now share this code with the user over email
	err := uls.notificationService.SendCodeOnEmail(email, code)
	if err != nil {
		return "", err
	}
	return code, nil
}

// ValidateLoginCodeAndInit this is the 2nd step in the login process and it is responsible for initializing
// defaults for the user if they do not exist
func (uls *userLoginService) ValidateLoginCodeAndInit(email, code string) (bool, error) {
	validated, err := uls.userService.ValidateEmailAndLoginCode(email, code)
	if err != nil {
		return false, err
	}
	if !validated {
		return false, nil
	}
	err = uls.initUser(email)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Wrapper around FindUserByEmail
func (uls *userLoginService) FindUserByEmail(email string) (*User, error) {
	return uls.userService.FindUserByEmail(email)
}

// initUser - Checks if there is a folder called Inbox present for this user, if not, create it
func (uls *userLoginService) initUser(email string) error {
	user, err := uls.userService.FindUserByEmail(email)
	if err != nil {
		return err
	}
	// Check if inbox folder is present
	foldersForUser, err := uls.folderService.GetAllByUserID(user.ID)
	// length will be empty if no records are found, but we can iterate and check if it has the Inbox folder
	for _, folder := range foldersForUser {
		if folder.Name == "Inbox" {
			// Return and do not do anything
			return nil
		}
	}
	// If we are here, it means that the folder does not exist
	// Create the Inbox folder
	inboxFolder := &Folder{Name: "Inbox", UserID: user.ID}
	_, err = uls.folderService.Create(inboxFolder)
	return err
}

type FourDigitCodeGenerator struct{}

func NewFourDigitCodeGenerator() CodeGenerator {
	return &FourDigitCodeGenerator{}
}
func (cg *FourDigitCodeGenerator) Generate() string {
	// We will generate a dummy code for now
	low := 1000
	high := 9999
	num := low + rand.Intn(high-low)
	return strconv.Itoa(num)
}
