package main

import (
	"fmt"
	"net/http"
	"net/smtp"
)

type sender struct {
	email    string
	passwrod string
}

type receiver struct {
	emails []string
}

func send(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(rw, r, "resources/send.html")
	case "POST":

		http.ServeFile(rw, r, "resources/send.html")
	}
}

func sendMail(senderData sender, receiverData receiver, message []byte) {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", senderData.email, senderData.passwrod, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderData.email, receiverData.emails, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message was send successfully.")
}

func setupRouts() {
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRouts()
}
