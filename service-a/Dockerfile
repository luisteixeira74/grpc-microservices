# Etapa 1: Build
FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app

# Copia go.mod e go.sum da raiz do projeto
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o projeto para dentro do container
COPY . .

# Compila o service-a (o main.go está dentro de service-a/)
RUN CGO_ENABLED=0 go build -o server ./service-a

# Etapa 2: Runtime
FROM alpine:latest
RUN adduser -D -g '' appuser
WORKDIR /app
COPY --from=builder /app/server .
RUN chmod +x server
USER appuser

CMD ["./server"]
