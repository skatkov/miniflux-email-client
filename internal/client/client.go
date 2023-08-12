package miniflux_email_client

import (
	"errors"
	"os"

	miniflux "miniflux.app/client"
)

type minifluxClient struct {
	miniflux    *miniflux.Client
	category_id int64
}

func NewClient() *minifluxClient {
	return &minifluxClient{
		miniflux: miniflux.New(os.Getenv("MINIFLUX_URL"), os.Getenv("MINIFLUX_TOKEN")),
	}
}

func (c *minifluxClient) GetUnreadEntries(category_name string) (*miniflux.EntryResultSet, error) {
	//TODO: we should support cases when category_name is not set.
	categories, err := c.miniflux.Categories()
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		if category.Title == category_name {
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
