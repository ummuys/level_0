FROM golang:1.24.5-alpine


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY cmd/producer/ ./cmd/producer/
COPY internal/testdata/ ./internal/testdata/


RUN go build -o /app/producer ./cmd/producer

ENTRYPOINT ["/app/producer"]