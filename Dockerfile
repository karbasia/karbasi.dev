# Build stage
FROM golang:1.24.4-alpine3.22 AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the API binary
RUN CGO_ENABLED=1 GOOS=linux go build -o api ./cmd/api

# Final stage
FROM alpine:3.22.0

WORKDIR /app

# Create goapi user and group
RUN addgroup -S goapi && adduser -S goapi -G goapi

USER goapi

# Copy the built binary from builder
COPY --from=builder /app/api /app/api

EXPOSE 8080

ENTRYPOINT ["/app/api"]
