package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	fmt.Println("Dima huima")

	from := "ebanyvrotblyatvashegocasino@gmail.com"
	password := "A123456789b"

	to := []string{
		"padrition@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("This is a message.")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
