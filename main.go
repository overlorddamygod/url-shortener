package main

import (
	"github.com/overlorddamygod/url-shortener/api"
)

func main() {
	const PORT = ":8080"
	server := api.NewServer()
	server.Run(PORT)
}
