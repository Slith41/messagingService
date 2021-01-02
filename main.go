package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
)

type sender struct {
	Email    string
	Password string
}

type receiver struct {
	Emails []string
}

func send(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		testJSON := `{"emails" : ["padrition@gmail.com", "dimastail23@gmail.com", "dmtevseev@gmail.com"]}`

		receiversEmails := parseEmailsInJSON(testJSON)

		var senderData sender
		senderData.Email = "ebanyvrotblyatvashegocasino@gmail.com"
		senderData.Password = "A123456789b"

		message := []byte("This is a robbery! Lay down and give me your money")

		sendMail(senderData, receiversEmails, message)

		http.ServeFile(rw, r, "resources/send.html")
	}
}

func sendMail(senderData sender, receiverData receiver, message []byte) {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", senderData.Email, senderData.Password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderData.Email, receiverData.Emails, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message was send successfully.")
}

func parseEmailsInJSON(JSONarray string) receiver {
	var receivers receiver
	json.Unmarshal([]byte(JSONarray), &receivers)

	return receivers
}
func setupRouts() {
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRouts()
}
