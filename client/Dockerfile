FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/client
RUN go build -o client

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/client/client .

USER nonroot:nonroot

CMD ["/app/client"]
