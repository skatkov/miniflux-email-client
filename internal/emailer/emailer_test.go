package emailer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	miniflux "miniflux.app/client"
)

func TestEmailerGetAdapter(t *testing.T) {
	emailer := NewEmailer("smtp.gmail.com", 587, "lunaticman@gmail.com", "testpass")

	assert.Equal(t, "smtp.gmail.com", emailer.GetAdapter().Server)
	assert.Equal(t, 587, emailer.GetAdapter().Port)
	assert.Equal(t, "lunaticman@gmail.com", emailer.GetAdapter().Username)
	assert.Equal(t, "testpass", emailer.GetAdapter().Password)
}

// entries mock
var entries miniflux.Entries = miniflux.Entries{
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
	emailer := NewEmailer("smtp.gmail.com", 587, "test@gmail.com", "testpass")

	assert.Equal(t, "test", emailer.GetMessage("newsletter@test.com", &ers))
}
