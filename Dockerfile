# Etapa 1: Compilação
FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN go build -o ./cmd/main ./cmd/main.go

# Etapa 2: Executar o binário
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/cmd/main .
CMD ["./main"]








