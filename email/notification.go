package email

import "github.com/naqvijafar91/focus"

type loginCodeSender struct {
}

func NewLoginCodeNotificationSender() focus.LoginCodeNotificationService {
	return &loginCodeSender{}
}

func (lcs *loginCodeSender) SendCodeOnEmail(email, code string) error {
	// Dummy implementation
	return nil
}
