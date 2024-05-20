package client

import (
	"errors"
	"log"

	miniflux "miniflux.app/client"
)

type Client struct {
	miniflux   *miniflux.Client
	categoryId int64
}

type MinifluxConfig struct {
	ApiUrl       string `env:"MINIFLUX_URL" envDefault:"https://reader.miniflux.app/"`
	Token        string `env:"MINIFLUX_TOKEN,required"`
	CategoryName string `env:"CATEGORY,required"`
}

func NewClient(config MinifluxConfig) *Client {
	log.Printf("Creating new client with config: %s", config)
	return &Client{
		miniflux: miniflux.New(config.ApiUrl, config.Token),
	}
}

// To do operation on categories, we need to retrieve category_id.
// Please note that category_id and category_name re different values, so we need to map appropriate category_id's.
//
// TODO: it's possible to set multiple categories, but support so far has not been tested.

func (c *Client) SetCategoryID(categoryName string) error {
	categories, err := c.miniflux.Categories()

	if err != nil {
		return errors.New("failed to find categories")
	}

	for _, category := range categories {
		if category.Title == categoryName {
			c.categoryId = category.ID
			log.Printf("category_id set to: %v", c.categoryId)
		}
	}
	return nil
}

// This returns unread entries, it requires currently that category_id has been set already with SetCategory method.
// TODO: We don't support retrieving unread entries without actegory_id, in reality it's rarely a good idea.
func (c *Client) GetUnreadEntries(limit int) (*miniflux.EntryResultSet, error) {
	if c.categoryId == 0 {
		// TODO: we should support cases when category_name is not set.
		return nil, errors.New("category_name is not set")
	}

	entries, err := c.miniflux.CategoryEntries(c.categoryId, &miniflux.Filter{
		Status:     miniflux.EntryStatusUnread,
		CategoryID: c.categoryId,
		Limit:      limit,
	})

	if err != nil {
		return nil, err
	}

	if entries.Total == 0 {
		return nil, errors.New("no unread entries found")
	}

	return entries, nil
}

// Marks entries as read, it requires currently that category_id has been set already with SetCategory method.
func (c *Client) MarkCategoryAsRead() error {
	return c.miniflux.MarkCategoryAsRead(c.categoryId)
}

func (c *Client) MarkAsRead(entries *miniflux.EntryResultSet) error {
	var entryIDs []int64
	for _, entry := range entries.Entries {
		entryIDs = append(entryIDs, entry.ID)
	}
	return c.miniflux.UpdateEntries(entryIDs, miniflux.EntryStatusRead)
}
