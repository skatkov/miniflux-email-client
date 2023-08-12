package miniflux_email_client

import (
	"errors"

	miniflux "miniflux.app/client"
)

type Client struct {
	miniflux    *miniflux.Client
	category_id int64
}

func NewClient(api_url string, token string) *Client {
	return &Client{
		miniflux: miniflux.New(api_url, token),
	}
}

func (c *Client) SetCategory(category_name string) error {
	categories, err := c.miniflux.Categories()
	if err != nil {
		return errors.New("failed to find categories")
	}

	for _, category := range categories {
		if category.Title == category_name {
			c.category_id = category.ID
		}
	}
	return nil
}

func (c *Client) GetUnreadEntries(category_name string) (*miniflux.EntryResultSet, error) {
	if c.category_id == 0 {
		//TODO: we should support cases when category_name is not set.
		return nil, errors.New("category_name is not set")
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

func (c *Client) MarkAsRead() error {
	return c.miniflux.MarkCategoryAsRead(c.category_id)
}
