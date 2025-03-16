# Use official Golang base image
FROM golang:1.21 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire backend source code
COPY . .

# Build the Golang application
RUN go build -o backend main.go

# Use a minimal runtime environment to reduce image size
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/backend .

# Expose application port
EXPOSE 8080

# Command to run the application
CMD ["./backend"]
