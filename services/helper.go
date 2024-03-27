package services

import (
	"fmt"
	"github.com/vaibhavgvk08/tigerhall-kittens/constants"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/common"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/mailer"
	"github.com/vaibhavgvk08/tigerhall-kittens/utils"
	"log"
)

func PrepareAndSendEmails(tiger *common.TigerDBStruct, reporterUsername string) {
	emailCh := make(chan common.EmailData, constants.MESSAGE_QUEUE_BUFFER_SIZE)
	// Removing the current user who has reported the tiger from email list. [Since reporter already knows the tiger location.]
	emailAddresList, _ := FetchUsersEmails(utils.RemoveAItemFromList(tiger.ReporterUserNamesList, reporterUsername))

	if len(emailAddresList) == 0 { //Possible when same user spots the tiger. In this case no other users have spotted the tiger, hence not sending the email.
		return
	}
	defer func() {
		close(emailCh)
	}()

	// Start the email sender goroutine
	go mailer.SendEmails(emailCh)

	// Produce email messages
	for i := 0; i < len(emailAddresList); i++ {
		email := common.CreateEmailObject(emailAddresList[i], tiger.Name, tiger.LastSeenTimeStamp[0], tiger.LastSeenCoordinates[0])

		emailCh <- email
		log.Println(fmt.Sprintf("Produced email %d.", i))
	}
}
