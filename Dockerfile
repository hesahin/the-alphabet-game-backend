FROM golang:1.22.2-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o build .

FROM debian:bullseye-slim

WORKDIR /app/cmd

COPY --from=builder /app/cmd/build .

CMD ["./build"]

EXPOSE 8080