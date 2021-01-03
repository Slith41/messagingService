package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"

	_ "github.com/lib/pq"
)

type Sender struct {
	Email    string
	Password string
}

type Receiver struct {
	Emails []string
}

type Email struct {
	email string
}

const (
	DB_USER     = "slith"
	DB_PASSWORD = "liac1912"
	DB_NAME     = "emails"
)

func DataFromDataBase() {

	connStr := "user=postgres password=liac1912 dbname=emails sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from emails")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	emails := []Email{}

	for rows.Next() {
		p := Email{}
		err := rows.Scan(&emails)
		if err != nil {
			fmt.Println(err)
			continue
		}
		emails = append(emails, p)
	}
	for _, p := range emails {
		fmt.Println(&p.email)
	}
}

func send(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		testJSON := `{"emails" : ["padrition@gmail.com", "dimastail23@gmail.com", "dmtevseev@gmail.com"]}`

		receiversEmails := parseEmailsInJSON(testJSON)

		var senderData Sender
		senderData.Email = "ebanyvrotblyatvashegocasino@gmail.com"
		senderData.Password = "A123456789b"

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
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message was send successfully.")
}

func parseEmailsInJSON(JSONarray string) Receiver {
	var receivers Receiver
	json.Unmarshal([]byte(JSONarray), &receivers)

	return receivers
}
func setupRouts() {
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {

	setupRouts()
	DataFromDataBase()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
