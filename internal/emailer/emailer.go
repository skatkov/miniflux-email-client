package emailer

import (
	"bytes"
	"fmt"
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
const smtp_address = "smtp.gmail.com:587"

var smtp_password = os.Getenv("GMAIL_PASSWORD")
var smtp_email = os.Getenv("GMAIL_EMAIL")

func NewEmailer() Emailer {
	return Emailer{
		adapter:      "gmail", // TODO: would be great to support more adapter, not just GMAIL SMTP.
		content_type: "html",  // TODO: should also support "text" as content type.
	}
}

func (e Emailer) Send(toEmail string, entries *miniflux.EntryResultSet) error {
	return sendEmail(toEmail, entries)
}

func auth() smtp.Auth {
	return smtp.PlainAuth("", smtp_email, smtp_password, smtp_server)
}

func sendEmail(toEmail string, entries *miniflux.EntryResultSet) error {
	msg := []byte("To: <" + toEmail + ">\r\n" +
		"Subject: " + subject() + "\r\n" +
		"Content-Type: text/html; charset=UTF-8" + "\r\n" +
		"\r\n" +
		formatBody(entries))

	return smtp.SendMail(smtp_address, auth(), smtp_email, []string{toEmail}, msg)
}

func subject() string {
	return fmt.Sprintf("📰 RSS Updates - %s", time.Now().Format("2006-01-02"))
}

func formatBody(entries *miniflux.EntryResultSet) string {
	var buffer bytes.Buffer

	for _, entry := range entries.Entries {
		buffer.WriteString(fmt.Sprintf("<h2><a href=\"%s\">%s</a></h2><br/>", entry.URL, entry.Title))
		buffer.WriteString(fmt.Sprintf("<div>%s</div>", entry.Content))
		buffer.WriteString("<hr>")
	}

	return buffer.String()
}
