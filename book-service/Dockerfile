# Stage 1: Build the Go binary
FROM golang:1.24.5 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o book-service main.go

# Stage 2: Minimal shell-capable base image
FROM alpine:3.18

# Install PostgreSQL client for pg_isready
RUN apk add --no-cache postgresql-client

WORKDIR /app

# Copy compiled Go binary and wait-for script
COPY --from=builder /app/book-service .
COPY wait-for.sh .

# Make script executable
RUN chmod +x wait-for.sh

# Expose port
EXPOSE 8080

# Run with wait-for DB
ENTRYPOINT ["./wait-for.sh", "db", "/app/book-service"]
