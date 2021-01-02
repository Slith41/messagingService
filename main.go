package main

import (
	"fmt"
	"net/http"
	"net/smtp"
)

func sendMail(rw http.ResponseWriter, r *http.Request) {

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	switch r.Method {
	case "GET":
		http.ServeFile(rw, r, "resources/send.html")
	case "POST":
		receiverEmail := []string{r.FormValue("receiverEmail")}
		message := []byte(r.FormValue("message"))
		senderEmail := r.FormValue("senderEmail")
		senderPassword := r.FormValue("senderPassword")
		auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, receiverEmail, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Email Sent Successfully!")
	}
}

func setupRouts() {
	http.HandleFunc("/send", sendMail)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRouts()
}
