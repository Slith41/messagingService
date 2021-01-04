package main

import "net/http"

func send(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		http.ServeFile(rw, r, "resources/send.html")
	}
}

func setupRouts() {
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}
