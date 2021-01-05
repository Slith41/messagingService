package main

import (
	"encoding/json"
	"net/http"
)

//Receiver struct contains an array of users to which a message will be send
type Receiver struct {
	Users []user
}

type user struct {
	Email string
}

func parseEmailsInJSON(r *http.Request) []string {
	var receivers Receiver
	emails := []string{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&receivers)
	checkErr(err)

	for _, user := range receivers.Users {
		emails = append(emails, user.Email)
	}

	return emails
}
