# Build stage
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

# Run stage (lightweight)
FROM ubuntu:22.04

WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/config ./config

EXPOSE 3000
CMD ["./app"]