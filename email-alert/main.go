package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	senderInfo := sender{email: "grainalert@gmail.com", password: "Grain123"}
	recieverInfo := reciever{email: "skyler.grainalert@gmail.com"}
	err := emailSetUpAndSend(senderInfo, recieverInfo)
	if err != nil {
		fmt.Printf("error sending email: %v", err)
	}
}

// func temp() {

// }

type sender struct {
	email    string
	password string
}

type reciever struct {
	email string
}

func emailSetUpAndSend(senderInfo sender, recieverInfo reciever) error {

	from := senderInfo.email
	password := senderInfo.password
	// I think that that it might be a good idea to create a seperate function that will check to make sure that the credential are they way we want them.
	// probably use a nested if statment.
	to := []string{
		recieverInfo.email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("This is a test email")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("email sent sucessfully!")
	return nil
}
