package main

import (
	"net/http"
)

func send(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		rw.Write([]byte(`Send here a POST request with json data`))

	case "POST":
		sender := Sender{Email: "**********", Password: "****"}
		sendMail(sender, parseEmailsInJSON(r), []byte(`TEST MESSAGE! IF YOU READ IT THAN JSON FILE IS PARSED AND EMAILS ARE OBTAINED CORECTLY!`))
	}
}

func setupRouts() {
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}
