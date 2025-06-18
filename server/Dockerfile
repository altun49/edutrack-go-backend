# Stage 1: build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/server

RUN go build -o /server

# Stage 2: minimal runtime
FROM debian:bookworm-slim

COPY --from=builder /server /server

ENTRYPOINT ["/server"]