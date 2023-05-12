package main

import (
	"net/smtp"
)

func Mailer(from string, to string, subject string, body string, password string, host string, port string) error{
	var auth = smtp.PlainAuth("", from, password, host)

	var err = smtp.SendMail(host + ":" + port, auth, from, []string{to}, []byte(body))

	return err
}
