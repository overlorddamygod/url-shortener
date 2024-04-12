package api

import (
	"github.com/gin-gonic/gin"
	"github.com/overlorddamygod/url-shortener/service"
)

type Server struct {
	router  *gin.Engine
	service *service.ShortenerService
}

func NewServer() *Server {
	return &Server{
		router:  gin.Default(),
		service: service.NewShortenerService(),
	}
}

func (s *Server) SetupRoutes() {
	s.router.POST("/shorten", s.ShortenUrlHandler)
	s.router.GET("/:shortCode", s.RedirectHandler)
}

func (s *Server) Run(port string) {
	s.SetupRoutes()
	s.router.Run(port)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Service() *service.ShortenerService {
	return s.service
}
