package client

import (
	"errors"
	"fmt"
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
	CategoryName string `env:"CATEGORY"`
	Limit        int    `env:"LIMIT"`
}

func (c MinifluxConfig) String() string {
	return fmt.Sprintf("ApiUrl: %s, CategoryName: %s, Limit: %d", c.ApiUrl, c.CategoryName, c.Limit)
}

func NewClient(config MinifluxConfig) *Client {
	log.Printf("Creating new client with config: %s", config)
	return &Client{
		miniflux: miniflux.New(config.ApiUrl, config.Token),
	}
}

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

	if c.categoryId == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (c *Client) GetUnreadEntries(limit int) (*miniflux.EntryResultSet, error) {
	entries, err := c.miniflux.Entries(&miniflux.Filter{
		Status:     miniflux.EntryStatusUnread,
		Limit:      limit,
		CategoryID: c.categoryId,
	})

	if err != nil {
		return nil, err
	}

	if entries.Total == 0 {
		return nil, errors.New("no unread entries found")
	}

	return entries, nil
}

func (c *Client) MarkAsRead(entries *miniflux.EntryResultSet) error {
	var entryIDs []int64
	for _, entry := range entries.Entries {
		entryIDs = append(entryIDs, entry.ID)
	}
	return c.miniflux.UpdateEntries(entryIDs, miniflux.EntryStatusRead)
}
