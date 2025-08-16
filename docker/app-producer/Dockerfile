FROM golang:1.24.5-alpine

# Рабочая папка
WORKDIR /app

# Сначала зависимости (лучше кешируются)
COPY go.mod go.sum ./
RUN go mod download

# Только нужный код
COPY cmd/producer/ ./cmd/producer/
COPY internal/testdata/ ./internal/testdata/

# Сборка бинаря
RUN go build -o /app/producer ./cmd/producer

ENTRYPOINT ["/app/producer"]