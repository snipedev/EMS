FROM golang:1.26.1-alpine AS builder

WORKDIR /app

# Ensure git is available to fetch modules and install ca-certificates for TLS
RUN apk add --no-cache git ca-certificates

# Take advantage of Docker layer caching by copying go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

COPY Incident-service/ .

ENV CGO_ENABLED=0 GOOS=linux

# Build a smaller static binary
RUN go build -ldflags="-s -w" -o main ./cmd

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]