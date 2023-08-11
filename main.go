package main

import (
	"fmt"
	"os"

	miniflux "github.com/skatkov/miniflux-email-client/internal/client"
	"github.com/skatkov/miniflux-email-client/internal/emailer"
)

var (
	receiverEmail = os.Getenv("RECEIVER_EMAIL")
)

func main() {
	client := miniflux.NewClient()
	mailer := emailer.NewEmailer(emailer.Plain)
	entries, err := client.GetUnreadEntries()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("sending email to: %s", receiverEmail)
	err = mailer.SendEmail(receiverEmail, entries)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.MarkAsRead()
	if err != nil {
		fmt.Println(err)
	}

}
