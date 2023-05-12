package mailer

import (
	"net/smtp"
)

func Mailer(from string, to string, subject string, body string, password string, host string, port string) {

	go func (){
		var auth = smtp.PlainAuth("", from, password, host)
	    smtp.SendMail(host + ":" + port, auth, from, []string{to}, []byte(body))
	}()
	
}
