#!/bin/bash

# Build script for URL Shortener Docker setup

echo "Building URL Shortener..."

# Build the Go binary for Linux
echo "Building Go binary..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-shortener main.go

# Build the Docker image
echo "Building Docker image..."
docker build -t url-shortener .

echo "Build complete!"
echo ""
echo "To run the container:"
echo "  docker run -p 8080:8080 url-shortener"
echo ""
echo "Or use docker-compose:"
echo "  docker compose up"