# Simple Dockerfile that uses a pre-built binary
# To build: 
# 1. CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-shortener main.go
# 2. docker build -t url-shortener .

FROM alpine:latest

WORKDIR /root/

# Copy the pre-built binary
COPY url-shortener .

# Make it executable
RUN chmod +x url-shortener

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./url-shortener"]