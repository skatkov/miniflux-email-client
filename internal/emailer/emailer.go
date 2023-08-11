package emailer

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"miniflux.app/client"
	miniflux "miniflux.app/client"
)

type AdapterInteface interface {
	SendEmail(string, *miniflux.EntryResultSet) error
	subject() string
	formatBody(*client.EntryResultSet) string
}

type MimeType string

const (
	HTML  MimeType = "text/html"
	Plain MimeType = "text/plain"
)

type EmailAdapter struct {
	content_type  MimeType
	smtp_server   string
	smtp_port     uint16
	smtp_email    string
	smtp_password string
}

func NewEmailer(adapter_name string, content_type MimeType) AdapterInteface {
	if adapter_name != "gmail" {
		// TODO: support more adapters, not just GMAIL.
		return nil
	}

	if len(content_type) > 0 {
		content_type = HTML
	}

	return &EmailAdapter{
		content_type:  content_type, // TODO: should also support "text" as content type.
		smtp_server:   "smtp.gmail.com",
		smtp_port:     587,
		smtp_password: os.Getenv("GMAIL_PASSWORD"),
		smtp_email:    os.Getenv("GMAIL_EMAIL"),
	}
}

func (a *EmailAdapter) auth() smtp.Auth {
	return smtp.PlainAuth("", a.smtp_email, a.smtp_password, a.smtp_server)
}

func (a *EmailAdapter) SendEmail(toEmail string, entries *miniflux.EntryResultSet) error {
	msg := []byte("To: <" + toEmail + ">\r\n" +
		"Subject: " + a.subject() + "\r\n" +
		"Content-Type: " + string(a.content_type) + "; charset=UTF-8" + "\r\n" +
		"\r\n" +
		a.formatBody(entries))
	address := fmt.Sprintf("%s:%d", a.smtp_server, a.smtp_port)

	return smtp.SendMail(address, a.auth(), a.smtp_email, []string{toEmail}, msg)
}

func (a *EmailAdapter) subject() string {
	return fmt.Sprintf("📰 RSS Updates - %s", time.Now().Format("2006-01-02"))
}

func (a *EmailAdapter) formatBody(entries *miniflux.EntryResultSet) string {
	var buffer bytes.Buffer

	for _, entry := range entries.Entries {
		buffer.WriteString(fmt.Sprintf("<h2><a href=\"%s\">%s</a></h2><br/>", entry.URL, entry.Title))
		buffer.WriteString(fmt.Sprintf("<div>%s</div>", entry.Content))
		buffer.WriteString("<hr>")
	}

	return buffer.String()
}
