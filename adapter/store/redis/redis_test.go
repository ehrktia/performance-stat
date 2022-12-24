package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_url(t *testing.T) {
	got := getURLFromEnv()
	assert.NotEmpty(t, got)
	assert.Contains(t, got, defaultHost)
	assert.Contains(t, got, defaultPort)
}
