package main

import (
	"fmt"
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
		fmt.Println("failed to fetch RSS updates: " + err.Error())
		return
	}

	fmt.Println("sending email to: %s", receiverEmail)
	err = mailer.SendEmail(receiverEmail, entries)

	if err != nil {
		fmt.Println("failed to send, due to an error: " + err.Error())
		return
	}

	err = client.MarkAsRead()
	if err != nil {
		fmt.Println("failed to mark RSS updates as read, due to an error: " + err.Error())
	}

}
