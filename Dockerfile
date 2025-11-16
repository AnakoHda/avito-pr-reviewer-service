FROM golang:1.25.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go mod tidy
RUN go build -o pr-service ./cmd/service/main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.25.0

FROM alpine

WORKDIR /app
COPY --from=builder /app/pr-service .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY migrations ./migrations

EXPOSE 8080

CMD ["./pr-service"]