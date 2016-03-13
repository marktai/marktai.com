package dummy

import (
	"email"
	"fmt"
	"log"
	"time"
)

func sendEmail() {

	mail := email.Email{}

	mail.Recipient = "taifighterm@gmail.com"
	mail.Subject = fmt.Sprintf("Test (%s)", time.Now().String())
	mail.Body = "this is a test email.\r\nIt has multiple lines"
	err := email.SendMail("mail.marktai.com:25", mail)
	if err != nil {
		log.Println(err)
	}

}

func Run() {
	// sendEmail()
}
