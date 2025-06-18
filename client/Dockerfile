# Stage 1: build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/client

RUN go build -o /client

# Stage 2: minimal runtime
FROM debian:bookworm-slim

COPY --from=builder /client /client

ENTRYPOINT ["/client"]