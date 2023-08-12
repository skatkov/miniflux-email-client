package emailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailer(t *testing.T) {
	emailer := NewEmailer("smtp.gmail.com", 587, "lunaticman@gmail.com", "testpass")

	assert.Equal(t, "smtp.gmail.com", emailer.GetAdapter().Server)
	assert.Equal(t, 587, emailer.GetAdapter().Port)
	assert.Equal(t, "lunaticman@gmail.com", emailer.GetAdapter().Username)
	assert.Equal(t, "testpass", emailer.GetAdapter().Password)
}
