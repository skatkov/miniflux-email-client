package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c := NewClient("http://localhost:8080", "token_test")

	err := c.SetCategory("test")

	assert.Error(t, err)
}
