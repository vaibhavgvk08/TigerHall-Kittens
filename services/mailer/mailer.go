package mailer

import (
	"fmt"
	"github.com/go-mail/mail"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/common"
	"log"
)

// email consumer
func SendEmails(emailCh <-chan common.EmailData) {
	for email := range emailCh {
		sendEmail(email)
	}
}

func sendEmail(emailData common.EmailData) {
	defer func() {
		log.Println("Email sent.\n")
	}()

	m := mail.NewMessage()

	m.SetHeader("From", EMAIL_SENDER)

	m.SetHeader("To", emailData.To)

	m.SetHeader("Subject", "TigerHall - Kittens!")

	m.SetBody("text/html", fmt.Sprintf(EMAIL_BODY, emailData.UserName, *emailData.TigerName, emailData.TigerLastSeenTimestamp, emailData.TigerLastSeenCoordinates.Lat, emailData.TigerLastSeenCoordinates.Long))

	d := mail.NewDialer("smtp.gmail.com", 587, "vaibhavgvk08@gmail.com", EMAIL_TOKEN)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
