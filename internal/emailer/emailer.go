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
	port         uint16
	email        string
	password     string
}

var (
	password string
	email    string
	server   string
)

func NewEmailer(content_type MimeType) AdapterInteface {
	if len(content_type) > 0 {
		content_type = TEXT
	}
	// TODO: we should dperecate usage of GMAIL_* env variables.
	if password = os.Getenv("SMTP_PASSWORD"); password == "" {
		password = os.Getenv("GMAIL_PASSWORD")
	}

	if email = os.Getenv("SMTP_EMAIL"); email == "" {
		email = os.Getenv("GMAIL_EMAIL")
	}

	if server := os.Getenv("SMTP_SERVER"); server == "" {
		server = "smtp.gmail.com"
	}

	return &SMTPAdapter{
		content_type: content_type,
		server:       server,
		port:         587, //TODO: should be possible to configure through ENV variable.
		password:     password,
		email:        email,
	}
}

func (a *SMTPAdapter) auth() smtp.Auth {
	return smtp.PlainAuth("", a.email, a.password, a.server)
}

func (a *SMTPAdapter) SendEmail(toEmail string, entries *miniflux.EntryResultSet) error {
	body, err := a.formatBody(entries)
	if err != nil {
		return err
	}
	msg := []byte("To: <" + toEmail + ">\r\n" +
		"Subject: " + a.subject() + "\r\n" +
		"Content-Type: " + string(a.content_type) + "; charset=UTF-8" +
		"\r\n" + body)

	return smtp.SendMail(a.server+":"+string(a.port), a.auth(), a.email, []string{toEmail}, msg)
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
			buffer.WriteString(fmt.Sprintf("%s - %s", entry.URL, entry.Title))
			buffer.WriteString(fmt.Sprintf("--------------\n%s\n--------------", entry.Content))
			buffer.WriteString("\n")
		}
	default:
		return "", errors.New("invalid content type - " + string(a.content_type))
	}

	return buffer.String(), nil
}
