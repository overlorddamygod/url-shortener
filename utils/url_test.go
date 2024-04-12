package utils_test

import (
	"testing"

	"github.com/overlorddamygod/url-shortener/utils"
	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	t.Run("Should recognize valid URLs", func(t *testing.T) {
		assert.True(t, utils.IsValidURL("https://www.google.com"))
		assert.True(t, utils.IsValidURL("https://www.instagram.com/eqwe/a.txt"))
		assert.True(t, utils.IsValidURL("https://www.google.com/search?q=golang"))
		assert.True(t, utils.IsValidURL("https://www.google.com/search?q=golang#top"))
		assert.True(t, utils.IsValidURL("http://localhost:8080"))
		assert.True(t, utils.IsValidURL("google.com"))
		assert.True(t, utils.IsValidURL("192.168.1.1"))
	})

	t.Run("Should recognize invalid URLs", func(t *testing.T) {
		assert.False(t, utils.IsValidURL(""))
		assert.False(t, utils.IsValidURL("invalid_url"))
		assert.False(t, utils.IsValidURL("https://.com"))
		assert.False(t, utils.IsValidURL("https://example.com/pa th"))
		assert.False(t, utils.IsValidURL("//example.com"))
		assert.False(t, utils.IsValidURL("http://example..com"))
	})
}
