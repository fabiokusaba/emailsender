package mail

import (
	"os"

	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
	"gopkg.in/gomail.v2"
)

func Send(campaign *campaign.Campaign) error {
	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	var emails []string
	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", campaign.Name)
	m.SetBody("text/html", campaign.Content)

	return d.DialAndSend(m)
}
