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
	mailer := emailer.NewEmailer()
	entries, err := client.GetUnreadEntries()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("sending email to: ", receiverEmail)
	err = mailer.Send(receiverEmail, entries)

	if err != nil {
		err = client.MarkAsRead()

		if err != nil {
			fmt.Println(err)
		}
	}
}
