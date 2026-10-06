# Build stage
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o gotask-api .

# Final lightweight image
FROM alpine:3.20

WORKDIR /app

# Copy compiled application
COPY --from=builder /app/gotask-api .

# Application port
EXPOSE 8080

# Start application
CMD ["./gotask-api"]