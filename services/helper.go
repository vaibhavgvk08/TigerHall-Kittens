package services

import (
	"fmt"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/common"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/mailer"
	"log"
)

func PrepareAndSendEmails(tiger *model.Tiger) {
	emailCh := make(chan common.EmailData, constants.MESSAGE_QUEUE_BUFFER_SIZE)
	emailAddresList := FetchUsersEmails(tiger.UsersWhoSightedTiger)

	defer func() {
		close(emailCh)
	}()

	// Start the email sender goroutine
	go mailer.SendEmails(emailCh)

	// Produce email messages
	for i := 0; i < len(emailAddresList); i++ {
		email := common.EmailData{
			To:                       emailAddresList[i],
			UserName:                 tiger.UsersWhoSightedTiger[i],
			TigerName:                tiger.Name,
			TigerLastSeenTimestamp:   tiger.LastSeenTimeStamp[0],
			TigerLastSeenCoordinates: tiger.LastSeenCoordinates[0],
		}
		emailCh <- email
		log.Println(fmt.Sprintf("Produced email %d\n", i))
	}
}
