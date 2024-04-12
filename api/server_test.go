package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/overlorddamygod/url-shortener/api"
	"github.com/stretchr/testify/assert"
)

func TestUrlShortenerServer(t *testing.T) {
	server := api.NewServer()
	server.SetupRoutes()

	router := server.Router()

	t.Run("[/shorten] Should throw error when no url is provided to request body", func(t *testing.T) {
		t.Parallel()
		payload := `{}`

		req, err := http.NewRequest(http.MethodPost, "/shorten", strings.NewReader(payload))

		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response api.ErrorResponse

		assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
		assert.Equal(t, "invalid request", response.Message)
	})

	t.Run("[/shorten] Should throw error for an invalid URL", func(t *testing.T) {
		t.Parallel()

		payload := `{"url": "invalid-url"}`

		req, err := http.NewRequest(http.MethodPost, "/shorten", strings.NewReader(payload))

		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response api.ErrorResponse

		assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
		assert.Equal(t, "invalid url", response.Message)
	})

	t.Run("[/shorten] Should successfully return a short url", func(t *testing.T) {
		t.Parallel()

		payload := `{"url": "https://example.com"}`

		req, err := http.NewRequest(http.MethodPost, "/shorten", strings.NewReader(payload))

		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response api.ShortenUrlResponse

		assert.NoError(t, json.NewDecoder(w.Body).Decode(&response))
		assert.NotEmpty(t, response.ShortUrl)
	})

	t.Run("[/{shortCode}] Should throw not found error", func(t *testing.T) {
		t.Parallel()

		randomCode := "U2XjkS"
		req, err := http.NewRequest(http.MethodGet, "/"+randomCode, nil)

		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("[/{shortCode}] Should redirect to original long url", func(t *testing.T) {
		t.Parallel()

		// Creating a short url for google.com
		const googleUrl = "https://google.com"
		googleShortCode, err := server.Service().GetShortCode(googleUrl)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "/"+googleShortCode, nil)

		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// check if it redirects to google.com
		assert.Equal(t, http.StatusMovedPermanently, w.Code)
		location := w.Header().Get("Location")
		assert.Equal(t, googleUrl, location)
	})
}
