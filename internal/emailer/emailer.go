package emailer

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	miniflux "miniflux.app/client"
)

type Emailer struct {
	adapter      string
	content_type string
}

const smtp_server = "smtp.gmail.com"

var smtp_password = os.Getenv("GMAIL_PASSWORD")
var smtp_email = os.Getenv("GMAIL_EMAIL")

func NewEmailer() Emailer {
	return Emailer{
		adapter:      "gmail",
		content_type: "html",
	}
}

func (e Emailer) Send(toEmail string, entries *miniflux.EntryResultSet) error {
	auth := smtp.PlainAuth("", smtp_email, smtp_password, smtp_server)
	subject := fmt.Sprintf("📰 News Updates - %s", time.Now().Format("2006-01-02"))

	msg := []byte("To: <" + toEmail + ">\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8" + "\r\n" +
		"\r\n" +
		e.formatBody(entries))

	err := smtp.SendMail("smtp.gmail.com:587", auth, smtp_email, []string{toEmail}, msg)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
		return err
	} else {
		log.Println("Email sent successfully")
		return nil
	}

}

func (emailer Emailer) formatBody(entries *miniflux.EntryResultSet) string {
	var buffer bytes.Buffer

	for _, entry := range entries.Entries {
		buffer.WriteString(fmt.Sprintf("<h2><a href=\"%s\">%s</a></h2><br/>", entry.URL, entry.Title))
		buffer.WriteString(fmt.Sprintf("<div>%s</div>", entry.Content))
		buffer.WriteString("<hr>")
	}

	return buffer.String()
}
