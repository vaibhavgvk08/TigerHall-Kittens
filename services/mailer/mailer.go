package mailer

import (
	"fmt"
	"github.com/go-mail/mail"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/common"
	"log"
	"os"
)

// email consumer
func SendEmails(emailCh <-chan common.EmailData) {
	for email := range emailCh {
		sendEmail(email)
	}
}

// todo - Integrate a better thrid party mailing service like mailtrap later.
func sendEmail(emailData common.EmailData) {
	defer func() {
		log.Println("Email sent.")
	}()

	m := mail.NewMessage()

	m.SetHeader("From", EMAIL_SENDER)

	m.SetHeader("To", emailData.To)

	m.SetHeader("Subject", EMAIL_SUBJECT)

	m.SetBody("text/html", fmt.Sprintf(EMAIL_BODY, *emailData.TigerName, emailData.TigerLastSeenTimestamp, emailData.TigerLastSeenCoordinates.Lat, emailData.TigerLastSeenCoordinates.Long))

	d := mail.NewDialer(EMAIL_HOST, 587, EMAIL_USERNAME, os.Getenv("password"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
