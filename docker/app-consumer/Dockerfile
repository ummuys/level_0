FROM golang:1.24.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/consumer/ ./cmd/consumer/
COPY internal/ ./internal/

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN /go/bin/swag init -g cmd/consumer/main.go
RUN go mod tidy 

RUN go build -o /app/consumer ./cmd/consumer


EXPOSE 8080
ENTRYPOINT ["/app/consumer"]
