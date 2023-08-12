package emailer

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
	"time"

	miniflux "miniflux.app/client"
)

type AdapterInteface interface {
	SendEmail(string, *miniflux.EntryResultSet, MimeType) error
	subject() string
	formatBody(*miniflux.EntryResultSet, MimeType) (string, error)
	Adapter() *SMTPAdapter
}

type MimeType string

const (
	HTML MimeType = "text/html"
	TEXT MimeType = "text/plain"
)

type SMTPAdapter struct {
	Server   string
	Port     int
	Username string
	Password string
}

func NewEmailer(server string, port int, username string, password string) AdapterInteface {
	return &SMTPAdapter{
		Server:   server,
		Port:     port,
		Password: password,
		Username: username,
	}
}

func (a *SMTPAdapter) auth() smtp.Auth {
	return smtp.PlainAuth("", a.Username, a.Password, a.Server)
}

func (a *SMTPAdapter) Adapter() *SMTPAdapter {
	return a
}

func (a *SMTPAdapter) SendEmail(toEmail string, entries *miniflux.EntryResultSet, contentType MimeType) error {
	if len(contentType) > 0 {
		contentType = TEXT
	}

	body, err := a.formatBody(entries, contentType)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("From: %s\r\n", a.Username)
	message += fmt.Sprintf("To: %s\r\n", []string{toEmail})
	message += fmt.Sprintf("Subject: %s\r\n", a.subject())
	message += fmt.Sprintf("Content-Type: %s; charset=UTF-8\r\n", string(contentType))
	message += fmt.Sprintf("\r\n%s\r\n", body)

	return smtp.SendMail(a.Server+":"+fmt.Sprint(a.Port), a.auth(), a.Username, []string{toEmail}, []byte(message))
}

func (a *SMTPAdapter) subject() string {
	return fmt.Sprintf("📰 RSS Updates - %s", time.Now().Format("2006-01-02"))
}

func (a *SMTPAdapter) formatBody(entries *miniflux.EntryResultSet, contentType MimeType) (string, error) {
	var buffer bytes.Buffer

	switch contentType {
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
		return "", errors.New("invalid content type - " + string(contentType))
	}

	return buffer.String(), nil
}
