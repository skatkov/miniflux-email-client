package main

import (
	"log"
	"os"
	"strconv"

	miniflux "github.com/skatkov/miniflux-email-client/internal/client"
	"github.com/skatkov/miniflux-email-client/internal/emailer"
)

var (
	sendTo        = os.Getenv("SEND_TO")
	categoryName  = os.Getenv("CATEGORY")
	minifluxUrl   = os.Getenv("MINIFLUX_URL")
	minifluxToken = os.Getenv("MINIFLUX_TOKEN")

	smtpServer   = os.Getenv("SMTP_SERVER")
	smtpUsername = os.Getenv("SMTP_USERNAME")
	smtpPassword = os.Getenv("SMTP_PASSWORD")
	smtpPort     int
	err          error
)

func main() {
	portStr := os.Getenv("SMTP_PORT")

	if portStr == "" {
		smtpPort = 587
	} else {
		smtpPort, err = strconv.Atoi(portStr)

		if err != nil {
			log.Fatalf("failed to parse port: %v", err)
		}

	}
	mailer := emailer.NewEmailer(smtpServer, smtpPort, smtpUsername, smtpPassword)
	client := miniflux.NewClient(minifluxUrl, minifluxToken)
	err = client.SetCategory(categoryName)

	if err != nil {
		log.Fatalf("failed to set category: %v", err)
	}

	entries, err := client.GetUnreadEntries(categoryName)

	if err != nil {
		log.Printf("failed to fetch RSS updates: %v", err)
		return
	}

	log.Printf("sending email to: %v", sendTo)
	err = mailer.Send(sendTo, entries)

	if err != nil {
		log.Fatalf("failed to send, due to an error: %v", err)
		return
	}

	err = client.MarkAsRead()
	if err != nil {
		log.Fatalf("failed to mark RSS updates as read, due to an error: %v", err)
	}

}
