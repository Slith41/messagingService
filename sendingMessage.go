package main

import (
	"fmt"
	"net/http"
	"net/smtp"
)

//Sender struct contains an information about the sender
type Sender struct {
	Email    string
	Password string
}

func send(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		receiversEmails := parseEmailsInJSON(emailDataJSON) //Parsing JSON to string for use as address of the recipient

		var senderData Sender
		senderData.Email = "***"
		senderData.Password = "***"

		message := []byte("This is a robbery! Lay down and give me your money")

		sendMail(senderData, receiversEmails, message)

		http.ServeFile(rw, r, "resources/send.html")
	}
}

func sendMail(senderData Sender, receiverData Receiver, message []byte) {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", senderData.Email, senderData.Password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderData.Email, receiverData.Emails, message)
	checkErr(err)

	fmt.Println("Message was send successfully.")
}
