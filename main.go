package main

import (
	"log"
	"os"
	"strconv"

	"github.com/caarlos0/env/v9"
	miniflux "github.com/skatkov/miniflux-email-client/internal/client"
	"github.com/skatkov/miniflux-email-client/internal/emailer"
)

var (
	sendTo = os.Getenv("SEND_TO")
	err    error
)

func main() {
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))

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

	log.Printf("setting category to: %v", minifluxConfig.CategoryName)
	err = client.SetCategoryID(minifluxConfig.CategoryName)

	if err != nil {
		log.Fatalf("failed to set category: %v", err)
	}

	entries, err := client.GetUnreadEntries(limit)

	if err != nil {
		log.Fatalf("failed to fetch RSS updates: %v", err)
		return
	}

	log.Printf("sending email to: %v", sendTo)
	err = mailer.Send(sendTo, entries)

	if err != nil {
		log.Fatalf("failed to send, due to an error: %v", err)
		return
	}

	if limit == 0 {
		err = client.MarkCategoryAsRead()
		if err != nil {
			log.Fatalf("failed to mark updates in category as read, due to an error: %v", err)
		}
	} else {
		err = client.MarkAsRead(entries)
		if err != nil {
			log.Fatalf("failed to mark individual RSS updates as read, due to an error: %v", err)
		}
	}

}
