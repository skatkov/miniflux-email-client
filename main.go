package main

import (
	"log"
	"os"

	"github.com/caarlos0/env/v9"
	miniflux "github.com/skatkov/miniflux-email-client/internal/client"
	"github.com/skatkov/miniflux-email-client/internal/emailer"
)

var (
	sendTo = os.Getenv("SEND_TO")

	err error
)

func main() {
	smtpConfig := emailer.SMTPConfig{}
	if err = env.Parse(&smtpConfig); err != nil {
		log.Fatalf("failed parsing ENV variables for SMTP: %+v\n", err)
	}

	mailer := emailer.NewEmailer(smtpConfig)

	minifluxConfig := miniflux.MinifluxConfig{}
	if err = env.Parse(&smtpConfig); err != nil {
		log.Fatalf("failed parsing ENV variables for miniflux: %+v\n", err)
	}
	client := miniflux.NewClient(minifluxConfig)
	err = client.SetCategoryID(minifluxConfig.CategoryName)

	if err != nil {
		log.Fatalf("failed to set category: %v", err)
	}

	entries, err := client.GetUnreadEntries()

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
