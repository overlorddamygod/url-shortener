package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/overlorddamygod/url-shortener/types"
)

func (s *Server) ShortenUrlHandler(c *gin.Context) {
	var req types.ShortenUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "invalid request",
		})
		return
	}

	shortCode, err := s.service.GetShortCode(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	shortUrl := protocol + "://" + c.Request.Host + "/" + shortCode

	c.JSON(http.StatusOK, types.ShortenUrlResponse{
		ShortUrl: shortUrl,
	})
}

func (s *Server) RedirectHandler(c *gin.Context) {
	shortCode := c.Param("shortCode")

	longUrl, err := s.service.GetLongUrl(shortCode)

	if err != nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longUrl)
}
