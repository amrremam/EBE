# Use official Golang image as base
FROM golang:1.23.2 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Set working directory to cmd/
WORKDIR /app/cmd

# Build the application
RUN go build -o main .

# Use a lightweight image for the final container
FROM debian:bullseye-slim

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/cmd/main .

# Expose the API port
EXPOSE 8080

# Run the application
CMD ["./main"]
