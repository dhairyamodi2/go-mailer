package main

import (
	"log"
	"net/smtp"
)

func Mailer(from string, to string, subject string, body string, password string, host string, port string) error{
	var auth = smtp.PlainAuth("", from, password, host)

	var err = smtp.SendMail(host + ":" + port, auth, from, []string{to}, []byte(body))

	return err
}
func main(){
	log.Print("Sending Mail...")
	var err = Mailer("ui20cs40@iiitsurat.ac.in", "dhairyam2002@gmail.com", "My subject", "Body", "ileoezejhktcjvoq", "smtp.gmail.com", "587")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Sent successfully!")
	}
}