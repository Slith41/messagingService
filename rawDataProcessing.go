package main

import "encoding/json"

//Receiver struct contains an array of users to which a message will be send
type Receiver struct {
	Users []user
}

type user struct {
	Email string
}

func parseEmailsInJSON(JSONarray string) []string {
	var receivers Receiver
	emails := []string{}

	json.Unmarshal([]byte(JSONarray), &receivers)

	for _, user := range receivers.Users {
		emails = append(emails, user.Email)
	}

	return emails
}
