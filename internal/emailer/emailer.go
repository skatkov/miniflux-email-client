package emailer

import (
	"bytes"
	"fmt"
	"net/smtp"
	"time"

	miniflux "miniflux.app/client"
)

type MimeType string

const (
	HTML MimeType = "text/html"
	TEXT MimeType = "text/plain"
)

type Emailer struct {
	ContentType MimeType
	SMTP        SMTPConfig
}

type SMTPConfig struct {
	Server   string
	Port     int
	Username string
	Password string
}

func NewEmailer(server string, port int, username string, password string) *Emailer {
	return &Emailer{
		ContentType: TEXT,
		SMTP: SMTPConfig{
			Server:   server,
			Port:     port,
			Password: password,
			Username: username,
		},
	}

}

func (e *Emailer) SetContentType(contentType MimeType) error {
	switch contentType {
	case HTML, TEXT:
		e.ContentType = contentType
		return nil
	default:
		return fmt.Errorf("invalid content type: %s", contentType)
	}
}

func (e *Emailer) Send(toEmail string, entries *miniflux.EntryResultSet) error {
	a := e.SMTP
	auth := smtp.PlainAuth("", a.Username, a.Password, a.Server)

	return smtp.SendMail(a.Server+":"+fmt.Sprint(a.Port), auth, a.Username, []string{toEmail}, []byte(e.getMessage(toEmail, entries)))
}

func (e *Emailer) getMessage(toEmail string, entries *miniflux.EntryResultSet) string {
	var body bytes.Buffer

	switch e.ContentType {
	case HTML:
		for _, entry := range entries.Entries {
			body.WriteString(fmt.Sprintf("<h2><a href=\"%s\">%s</a></h2><br/>", entry.URL, entry.Title))
			body.WriteString(fmt.Sprintf("<div>%s</div>", entry.Content))
			body.WriteString("<hr>")
		}
	case TEXT:
		for _, entry := range entries.Entries {
			body.WriteString(fmt.Sprintf("%s\n %s \n", entry.Title, entry.URL))
			body.WriteString("---\n")
		}
	}

	message := fmt.Sprintf("To: %s\r\n", []string{toEmail})
	message += fmt.Sprintf("Subject: %s\r\n", fmt.Sprintf("📰 RSS Updates - %s", time.Now().Format("2006-01-02")))
	message += fmt.Sprintf("Content-Type: %s; charset=UTF-8\r\n", e.ContentType)
	message += fmt.Sprintf("\r\n%s\r\n", body.String())

	return message
}
