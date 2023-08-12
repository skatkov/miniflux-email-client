package emailer

import (
	"bytes"
	"fmt"
	"net/smtp"
	"time"

	miniflux "miniflux.app/client"
)

type AdapterInteface interface {
	Send(string, *miniflux.EntryResultSet) error
	GetMessage(string, *miniflux.EntryResultSet) string
	GetAdapter() *SMTPAdapter
	SetContentType(MimeType) error
}

type MimeType string

const (
	HTML MimeType = "text/html"
	TEXT MimeType = "text/plain"
)

type Emailer struct {
	ContentType MimeType
	Adapter     SMTPAdapter
}

type SMTPAdapter struct {
	Server   string
	Port     int
	Username string
	Password string
}

func NewEmailer(server string, port int, username string, password string) AdapterInteface {
	return &Emailer{
		ContentType: TEXT,
		Adapter: SMTPAdapter{
			Server:   server,
			Port:     port,
			Password: password,
			Username: username,
		},
	}

}

func (e *Emailer) GetAdapter() *SMTPAdapter {
	return &e.Adapter
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
	a := e.GetAdapter()
	auth := smtp.PlainAuth("", a.Username, a.Password, a.Server)

	return smtp.SendMail(a.Server+":"+fmt.Sprint(a.Port), auth, a.Username, []string{toEmail}, []byte(e.GetMessage(toEmail, entries)))
}

func (e *Emailer) GetMessage(toEmail string, entries *miniflux.EntryResultSet) string {
	a := e.GetAdapter()

	message := fmt.Sprintf("From: %s\r\n", a.Username)
	message += fmt.Sprintf("To: %s\r\n", []string{toEmail})
	message += fmt.Sprintf("Subject: %s\r\n", e.GetSubject())
	message += fmt.Sprintf("Content-Type: %s; charset=UTF-8\r\n", string(e.ContentType))
	message += fmt.Sprintf("\r\n%s\r\n", e.GetBody(entries))

	return message
}

func (e *Emailer) GetSubject() string {
	return fmt.Sprintf("📰 RSS Updates - %s", time.Now().Format("2006-01-02"))
}

func (e *Emailer) GetBody(entries *miniflux.EntryResultSet) string {
	var buffer bytes.Buffer

	switch e.ContentType {
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
	}

	return buffer.String()
}
