package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	config MinifluxConfig = MinifluxConfig{
		ApiUrl:       "https://reader.miniflux.app/",
		Token:        "test01",
		CategoryName: "['daily']",
	}
)

func TestClient(t *testing.T) {
	c := NewClient(config)
	err := c.SetCategoryID(config.CategoryName)

	assert.Error(t, err)
}
