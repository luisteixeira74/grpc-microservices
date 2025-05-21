FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Copia todo o contexto do projeto
COPY . .

RUN go mod download

WORKDIR /app/service-b

RUN CGO_ENABLED=0 go build -o main .

FROM alpine:latest

RUN adduser -D -g '' appuser

WORKDIR /app/service-b

COPY --from=builder /app/service-b/main .

RUN chmod +x main

USER appuser

CMD ["./main"]
