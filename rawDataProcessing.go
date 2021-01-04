package main

import "encoding/json"

//Receiver struct contains an array of emails to which a message will be send
type Receiver struct {
	Emails []string
}

func parseEmailsInJSON(JSONarray string) Receiver {
	var receivers Receiver
	json.Unmarshal([]byte(JSONarray), &receivers)

	return receivers
}
