package main

import (
	"fmt"
	"net/smtp"
)

//Sender struct contains an information about the sender
type Sender struct {
	Email    string
	Password string
}

func sendMail(senderData Sender, receiversEmails []string, message []byte) {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", senderData.Email, senderData.Password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderData.Email, receiversEmails, message)
	checkErr(err)

	fmt.Println("Message was send successfully.")
}
