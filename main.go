package main

import (
	"time"

	_ "github.com/lib/pq"
)

func main() {
	timeNow := time.Now()
	userData := User{email: "*****", registrationStatus: false, timeSendingRegistrationMessage: timeNow}
	eventHandler(userData, 2)
	setupRouts()

}
