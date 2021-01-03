package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	_ "github.com/lib/pq"
)

type dbinfo struct {
	dbDriver   string
	dbUser     string
	dbPassword string
	dbName     string
}

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

		testJSON := `{"emails" : [***]}`

		receiversEmails := parseEmailsInJSON(testJSON)

		var senderData sender
		senderData.Email = "***"
		senderData.Password = "***"

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

func insertEmailIntoTable(db dbinfo, table string, email string) {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db.dbUser, db.dbPassword, db.dbName)
	database, err := sql.Open(db.dbDriver, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	err = database.QueryRow("INSERT INTO "+table+"(email) VALUES($1);", email).Scan()
	if err != nil {
		fmt.Println(err)
	}
}

func insertEmailsIntoTable(db dbinfo, table string, emails []string) {
	for _, email := range emails {
		insertEmailIntoTable(db, table, email)
	}
}
func selectAllFromTable(db dbinfo, table string) map[string]time.Time {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db.dbUser, db.dbPassword, db.dbName)
	database, err := sql.Open(db.dbDriver, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := database.Query("SELECT * FROM " + table + ";")
	if err != nil {
		fmt.Println(err)
	}
	var emailsMap map[string]time.Time
	for rows.Next() {
		var email string
		var createdAt time.Time
		err = rows.Scan(&email, &createdAt)
		if err != nil {
			fmt.Println(err)
		}
		emailsMap[email] = createdAt
	}
	return emailsMap
}
func setupRouts() {
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRouts()
}
