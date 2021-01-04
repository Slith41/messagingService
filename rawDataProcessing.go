package main

import "encoding/json"

//Receiver struct contains an array of users to which a message will be send
type Receiver struct {
	Users []user
}

type user struct {
	Email string
}

func parseEmailsInJSON(JSONarray string) Receiver {
	var receivers Receiver
	json.Unmarshal([]byte(JSONarray), &receivers)

	return receivers
}
