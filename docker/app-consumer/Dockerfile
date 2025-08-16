FROM golang:1.24.5-alpine

# Рабочая папка
WORKDIR /app


# Сначала зависимости (лучше кешируются)
COPY go.mod go.sum ./
RUN go mod download


# Только нужный код
COPY cmd/consumer/ ./cmd/consumer/
COPY internal/ ./internal/

RUN apk add --no-cache git && go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/consumer/main.go

# Сборка бинаря
RUN go build -o /app/consumer ./cmd/consumer

EXPOSE 8080
ENTRYPOINT ["/app/consumer"]
