package main

import (
	"log"
	"os"

	miniflux "github.com/skatkov/miniflux-email-client/internal/client"
	"github.com/skatkov/miniflux-email-client/internal/emailer"
)

var (
	receiverEmail = os.Getenv("RECEIVER_EMAIL")
	category_name = os.Getenv("CATEGORY")
)

func main() {
	client := miniflux.NewClient()
	mailer := emailer.NewEmailer(emailer.TEXT)
	entries, err := client.GetUnreadEntries(category_name)

	if err != nil {
		log.Fatalf("failed to fetch RSS updates: %v", err)
		return
	}

	log.Printf("sending email to: %v", receiverEmail)
	err = mailer.SendEmail(receiverEmail, entries)

	if err != nil {
		log.Fatalf("failed to send, due to an error: %v", err)
		return
	}

	err = client.MarkAsRead()
	if err != nil {
		log.Fatalf("failed to mark RSS updates as read, due to an error: %v", err)
	}

}
