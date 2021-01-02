package main

import (
	"fmt"
	"net/http"
	"net/smtp"
)

func sendMail(rw http.ResponseWriter, r *http.Request) {
	from := "ebanyvrotblyatvashegocasino@gmail.com"
	password := "A123456789b"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	switch r.Method {
	case "GET":
		http.ServeFile(rw, r, "resources/send.html")
	case "POST":
		receiverEmail := []string{r.FormValue("receiverEmail")}
		message := []byte(r.FormValue("message"))
		auth := smtp.PlainAuth("", from, password, smtpHost)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, receiverEmail, message)
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
