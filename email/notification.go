package email

import (
	"bytes"
	"errors"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"

	"github.com/naqvijafar91/focus"
)

type loginCodeSender struct {
	mailer *gmailSender
}

func NewLoginCodeNotificationSender(email, password string) (focus.LoginCodeNotificationService, error) {
	if strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
		return nil, errors.New("Invalid email or password")
	}
	return &loginCodeSender{newSender(email, password)}, nil
}

func (lcs *loginCodeSender) SendCodeOnEmail(email, code string) error {
	text := fmt.Sprintf("Your login code is %s", code)
	subject := "Your Focus App Login Code"
	messageBody := lcs.mailer.WritePlainEmail([]string{email}, subject, text)
	return lcs.mailer.SendMail([]string{email}, subject, messageBody)
}

const (
	/**
		Gmail SMTP Server
	**/
	SMTPServer = "smtp.gmail.com"
)

type gmailSender struct {
	User     string
	Password string
}

func newSender(Username, Password string) *gmailSender {

	return &gmailSender{Username, Password}
}

func (sender gmailSender) SendMail(Dest []string, Subject, bodyMessage string) error {

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {
		return fmt.Errorf("smtp error: %s", err)
	}
	return nil
}

func (sender gmailSender) WriteEmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *gmailSender) WriteHTMLEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/html", subject, bodyMessage)
}

func (sender *gmailSender) WritePlainEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/plain", subject, bodyMessage)
}
