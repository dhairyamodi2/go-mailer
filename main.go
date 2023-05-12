package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func Mailer(from string, to string, subject string, body string, password string, host string, port string, wg *sync.WaitGroup) {

	// go func (){
		var auth = smtp.PlainAuth("", from, password, host)
		defer wg.Done()
		var err = smtp.SendMail(host + ":" + port, auth, from, []string{to}, []byte(body))
		log.Print(err)
	// }()
}


func main(){
	fmt.Print(time.Now())
	var wg sync.WaitGroup
	for i := 1; i < 20; i++ {
		wg.Add(1)
		Mailer("ui20cs40@iiitsurat.ac.in", "dhairyam2811@gmail.com", "My subject", "Message", "ileoezejhktcjvoq", "smtp.gmail.com", "587", &wg)	
	}
	wg.Wait()
	fmt.Print(time.Now())
	
}