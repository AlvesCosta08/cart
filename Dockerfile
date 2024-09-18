# Etapa 1: Construir o binário
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Etapa 2: Executar o binário
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]

