FROM golang:1.22.6-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY . .

RUN go mod download
RUN go build -o api-server ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-server .

EXPOSE 8080

CMD ["./api-server-cep"]