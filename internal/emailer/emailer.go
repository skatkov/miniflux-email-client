package emailer

import (
	"bytes"
	"errors"
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
	formatBody(*client.EntryResultSet) (string, error)
}

type MimeType string

const (
	HTML MimeType = "text/html"
	TEXT MimeType = "text/plain"
)

type SMTPAdapter struct {
	content_type MimeType
	server       string
	port         string // TODO: "uint16" is better suited here.
	username     string
	password     string
}

func NewEmailer(content_type MimeType) AdapterInteface {
	if len(content_type) > 0 {
		content_type = TEXT
	}

	var port string
	if port := os.Getenv("SMTP_PORT"); len(port) > 0 {
		port = "587"
	}

	return &SMTPAdapter{
		content_type: content_type,
		server:       os.Getenv("SMTP_SERVER"),
		port:         port,
		password:     os.Getenv("SMTP_PASSWORD"),
		username:     os.Getenv("SMTP_USERNAME"),
	}
}

func (a *SMTPAdapter) auth() smtp.Auth {
	return smtp.PlainAuth("", a.username, a.password, a.server)
}

func (a *SMTPAdapter) SendEmail(toEmail string, entries *miniflux.EntryResultSet) error {
	body, err := a.formatBody(entries)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("From: %s\r\n", a.username)
	message += fmt.Sprintf("To: %s\r\n", []string{toEmail})
	message += fmt.Sprintf("Subject: %s\r\n", a.subject())
	message += fmt.Sprintf("Content-Type: %s; charset=UTF-8\r\n", string(a.content_type))
	message += fmt.Sprintf("\r\n%s\r\n", body)

	return smtp.SendMail(a.server+":"+a.port, a.auth(), a.username, []string{toEmail}, []byte(message))
}

func (a *SMTPAdapter) subject() string {
	return fmt.Sprintf("📰 RSS Updates - %s", time.Now().Format("2006-01-02"))
}

func (a *SMTPAdapter) formatBody(entries *miniflux.EntryResultSet) (string, error) {
	var buffer bytes.Buffer

	switch a.content_type {
	case HTML:
		for _, entry := range entries.Entries {
			buffer.WriteString(fmt.Sprintf("<h2><a href=\"%s\">%s</a></h2><br/>", entry.URL, entry.Title))
			buffer.WriteString(fmt.Sprintf("<div>%s</div>", entry.Content))
			buffer.WriteString("<hr>")
		}
	case TEXT:
		for _, entry := range entries.Entries {

			buffer.WriteString(fmt.Sprintf("%s\n %s \n", entry.Title, entry.URL))
			buffer.WriteString("---\n")
		}
	default:
		return "", errors.New("invalid content type - " + string(a.content_type))
	}

	return buffer.String(), nil
}
