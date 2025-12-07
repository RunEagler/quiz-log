FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install tools and generate code
RUN make install
RUN make generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /quiz-log-server .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS and sql-migrate for migrations
RUN apk --no-cache add ca-certificates postgresql-client

# Copy the binary from builder
COPY --from=builder /quiz-log-server /app/server

# Copy migration files and config
COPY db/migrations ./db/migrations
COPY dbconfig.yml ./

# Install sql-migrate
COPY --from=builder /go/bin/sql-migrate /usr/local/bin/

# Copy entrypoint script
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]
