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
	err    error
)

func main() {
	smtpConfig := emailer.SMTPConfig{}
	if err = env.Parse(&smtpConfig); err != nil {
		log.Fatalf("failed parsing ENV variables for SMTP: %+v\n", err)
	}

	mailer := emailer.NewEmailer(smtpConfig, emailer.HTML)

	minifluxConfig := miniflux.MinifluxConfig{}
	if err = env.Parse(&minifluxConfig); err != nil {
		log.Fatalf("failed parsing ENV variables for miniflux: %+v\n", err)
	}
	client := miniflux.NewClient(minifluxConfig)

	if minifluxConfig.CategoryName != "" {
		log.Printf("categoryName set to: %v", minifluxConfig.CategoryName)
		err = client.SetCategoryID(minifluxConfig.CategoryName)

		if err != nil {
			log.Fatalf("failed to set categoryId: %v", err)
			return
		}
	} else {
		log.Printf("categoryName is not set, fetching all entries")
	}
	entries, err := client.GetUnreadEntries(minifluxConfig.Limit)

	if err != nil {
		log.Fatalf("failed to fetch RSS updates: %v", err)
		return
	}

	log.Printf("sending email to: %v", sendTo)
	err = mailer.Send(sendTo, &minifluxConfig.CategoryName, entries)

	if err != nil {
		log.Fatalf("failed to send, due to an error: %v", err)
		return
	}

	err = client.MarkAsRead(entries)
	if err != nil {
		log.Fatalf("failed to mark RSS updates as read, due to an error: %v", err)
	}
}
