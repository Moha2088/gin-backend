# Start from the official Golang image for building the app
FROM golang:latest AS builder

WORKDIR /app

# Install git for go mod if needed
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose port (default Gin port)
EXPOSE 8081

# Run the binary
CMD ["./main"]