package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message" binding:"required"`
}

type ShortenUrlRequest struct {
	URL string `json:"url" binding:"required"`
}

type ShortenUrlResponse struct {
	ShortUrl string `json:"short_url" binding:"required"`
}

func (s *Server) ShortenUrlHandler(c *gin.Context) {
	var req ShortenUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "invalid request",
		})
		return
	}

	shortCode, err := s.service.GetShortCode(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	shortUrl := protocol + "://" + c.Request.Host + "/" + shortCode

	c.JSON(http.StatusOK, ShortenUrlResponse{
		ShortUrl: shortUrl,
	})
}

func (s *Server) RedirectHandler(c *gin.Context) {
	shortCode := c.Param("shortCode")

	longUrl, err := s.service.GetLongUrl(shortCode)

	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longUrl)
}
