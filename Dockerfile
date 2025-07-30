# Stage 1: Build Stage
FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install Air for live-reloading (optional, if you need it)
RUN go install github.com/air-verse/air@latest

# Copy source code
COPY . .

# Build the application with optimization flags
RUN go build -o app -ldflags="-s -w"

# Stage 2: Runtime Stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata
RUN mkdir -p /var/log/api/

# Set working directory
WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/app .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./app"]