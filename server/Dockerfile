FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/server
RUN go build -o server

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/server/server .

USER nonroot:nonroot

CMD ["/app/server"]
