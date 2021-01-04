package main

import "encoding/json"

type Receiver struct {
	Emails []string
}

func parseEmailsInJSON(JSONarray string) Receiver {
	var receivers Receiver
	json.Unmarshal([]byte(JSONarray), &receivers)

	return receivers
}
