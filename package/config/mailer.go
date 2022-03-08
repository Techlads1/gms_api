package config

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type Mail struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewMail(to []string, subject, body string) *Mail {

	return &Mail{
		to:      to,
		subject: subject,
		body:    body,
	}

}

func (mail *Mail) SendEmail() (bool, error) {

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + mail.subject + "!\n"
	body := []byte(subject + mime + "\n" + mail.body)

	if err := smtp.SendMail(GetMailSMTPAddress(), GetMailSMTPAuthentication(), "", mail.to, body); err != nil {
		return false, err
	}

	return true, nil
}

func (mail *Mail) ParseMailTemplate(templateMailFileName string, data interface{}) error {

	t, err := template.ParseFiles(templateMailFileName)

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return err
	}

	mail.body = buf.String()

	return nil
}