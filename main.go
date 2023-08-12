package main

import (
	"log"
	"os"
	"strconv"

	miniflux "github.com/skatkov/miniflux-email-client/internal/client"
	"github.com/skatkov/miniflux-email-client/internal/emailer"
)

var (
	receiverEmail = os.Getenv("RECEIVER_EMAIL")
	category_name = os.Getenv("CATEGORY")
	smtp_server   = os.Getenv("SMTP_SERVER")
	smtp_username = os.Getenv("SMTP_USERNAME")
	smtp_password = os.Getenv("SMTP_PASSWORD")
	smtp_port     int
	err           error
)

func main() {
	portStr := os.Getenv("SMTP_PORT")

	if portStr == "" {
		smtp_port = 587
	} else {
		smtp_port, err = strconv.Atoi(portStr)

		if err != nil {
			log.Fatalf("failed to parse port: %v", err)
		}

	}

	client := miniflux.NewClient()
	mailer := emailer.NewEmailer(smtp_server, smtp_port, smtp_username, smtp_password)
	entries, err := client.GetUnreadEntries(category_name)

	if err != nil {
		log.Printf("failed to fetch RSS updates: %v", err)
		return
	}

	log.Printf("sending email to: %v", receiverEmail)
	err = mailer.Send(receiverEmail, entries)

	if err != nil {
		log.Fatalf("failed to send, due to an error: %v", err)
		return
	}

	err = client.MarkAsRead()
	if err != nil {
		log.Fatalf("failed to mark RSS updates as read, due to an error: %v", err)
	}

}
