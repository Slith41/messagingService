package main

import "time"

type User struct {
	email                          string
	registrationStatus             bool
	timeSendingRegistrationMessage time.Time
}

func eventHandler(userData User, timePeriod int) {
	timePause := time.NewTimer(time.Duration(timePeriod) * time.Minute)
	sender := Sender{Email: "*******", Password: "A123456789b"}
	resieversEmails := []string{userData.email}
	message := []byte(`hello`)
	<-timePause.C
	if userData.registrationStatus == false {

		sendMail(sender, resieversEmails, message)
	}

}
