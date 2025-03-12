FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/bot cmd/main.go

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /go/bin/bot /app/bot

EXPOSE 8080

CMD ["/app/bot"]