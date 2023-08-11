package miniflux_email_client

import (
	"errors"
	"os"

	miniflux "miniflux.app/client"
)

type minifluxClient struct {
	miniflux      *miniflux.Client
	category_name string
	category_id   int64
}

func NewClient() *minifluxClient {
	return &minifluxClient{
		miniflux:      miniflux.New(os.Getenv("MINIFLUX_URL"), os.Getenv("MINIFLUX_TOKEN")),
		category_name: os.Getenv("CATEGORY"),
	}
}

func (c *minifluxClient) GetUnreadEntries() (*miniflux.EntryResultSet, error) {
	categories, err := c.miniflux.Categories()
	if err != nil {
		return nil, err
	}

	// print out category that has daily value
	for _, category := range categories {
		if category.Title == c.category_name {
			c.category_id = category.ID
		}
	}

	entries, err := c.miniflux.CategoryEntries(c.category_id, &miniflux.Filter{Status: miniflux.EntryStatusUnread, CategoryID: c.category_id})
	if err != nil {
		return nil, err
	}

	if entries.Total == 0 {
		return nil, errors.New("no unread entries found")
	}

	return entries, nil
}

func (c *minifluxClient) MarkAsRead() error {
	return c.miniflux.MarkCategoryAsRead(c.category_id)

}
