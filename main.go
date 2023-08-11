package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skatkov/miniflux-email-client/internal/emailer"
	miniflux "miniflux.app/client"
)

type Entry struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

type EntriesResponse struct {
	Entries []Entry `json:"entries"`
}

func main() {
	minifluxURL := os.Getenv("MINIFLUX_URL")
	token := os.Getenv("MINIFLUX_TOKEN")
	receiverEmail := os.Getenv("RECEIVER_EMAIL")
	selected_category := os.Getenv("CATEGORY")

	client := miniflux.New(minifluxURL, token)

	categories, err := client.Categories()
	if err != nil {
		fmt.Println(err)
	}

	var category_id int64

	for _, category := range categories {
		if category.Title == selected_category {
			category_id = category.ID
		}
	}

	fmt.Printf("Detected category_id is %d", category_id)

	entries, err := client.CategoryEntries(category_id, &miniflux.Filter{Status: miniflux.EntryStatusUnread, CategoryID: category_id})
	if err != nil {
		fmt.Println(err)
	}

	if entries.Total == 0 {
		log.Println("No unread entries found")
		return
	}

	mailer := emailer.NewEmailer()

	fmt.Println("sending email to: ", receiverEmail)
	err = mailer.Send(receiverEmail, entries)

	if err != nil {
		err = client.MarkCategoryAsRead(category_id)
		if err != nil {
			fmt.Println(err)
		}
	}

}
