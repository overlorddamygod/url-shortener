package types

type ErrorResponse struct {
	Message string `json:"message" binding:"required"`
}

type ShortenUrlRequest struct {
	URL string `json:"url" binding:"required"`
}

type ShortenUrlResponse struct {
	ShortUrl string `json:"short_url" binding:"required"`
}
