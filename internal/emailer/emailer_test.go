package emailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	config SMTPConfig = SMTPConfig{
		Server:   "smtp.gmail.com",
		Port:     587,
		Username: "test@test.com",
		Password: "testpass",
	}
)

func TestEmailerGetAdapter(t *testing.T) {
	emailer := NewEmailer(config, HTML)

	assert.Equal(t, "smtp.gmail.com", emailer.SMTP.Server)
	assert.Equal(t, 587, emailer.SMTP.Port)
	assert.Equal(t, "test@test.com", emailer.SMTP.Username)
	assert.Equal(t, "testpass", emailer.SMTP.Password)
}

func TestNewEmailerFromFallback(t *testing.T) {
	cfg := SMTPConfig{
		Server:   "smtp.gmail.com",
		Port:     587,
		Username: "user@example.com",
		Password: "pass",
	}
	emailer := NewEmailer(cfg, HTML)

	assert.Equal(t, "user@example.com", emailer.SMTP.From)
}

func TestNewEmailerFromExplicit(t *testing.T) {
	cfg := SMTPConfig{
		Server:   "smtp.gmail.com",
		Port:     587,
		Username: "user@example.com",
		Password: "pass",
		From:     "sender@example.com",
	}
	emailer := NewEmailer(cfg, HTML)

	assert.Equal(t, "sender@example.com", emailer.SMTP.From)
}

// entries mock
/* var entries miniflux.Entries = miniflux.Entries{
	&miniflux.Entry{
		ID:          1,
		UserID:      1,
		FeedID:      1,
		Status:      "new",
		Hash:        "hash",
		Title:       "entry title",
		URL:         "http://www.example.com/news1",
		CommentsURL: "http://www.example.com/news1/comments",
		Date:        time.Now(),
		CreatedAt:   time.Now(),
		ChangedAt:   time.Now(),
		Content:     "entry content",
		Author:      "entry author",
		ShareCode:   "share code",
		Starred:     false,
		ReadingTime: 5,
		Feed:        &miniflux.Feed{ID: 1},
		Tags:        []string{"tag1", "tag2"},
	},
}

// EntryResultSet mock
var ers = miniflux.EntryResultSet{
	Total:   len(entries),
	Entries: entries,
}

func TestEmilerGetMessage(t *testing.T) {
	emailer := NewEmailer(config)

	//assert.Equal(t, "test", emailer.getMessage("newsletter@test.com", &ers))
}
*/
