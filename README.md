# grpc-microservices

Projeto exemplo de microserviços em Go comunicando via gRPC, com Docker e Docker Compose.

---

## Estrutura

- `service-a/` — Serviço gRPC servidor que gera uma stream de palavras aleatórias.
- `service-b/` — Cliente gRPC que consome a stream do `service-a`.
- `proto/` — Arquivos `.proto` com as definições das mensagens e serviços gRPC.
- `docker-compose.yml` — Orquestra os containers `service-a` e `service-b`.
- `Dockerfile` — Dockerfiles para build de cada serviço.
- `go.mod` e `go.sum` — Módulos e dependências Go comuns a ambos os serviços.

---

## Requisitos

- Docker (versão recente)
- Docker Compose
- Protocol Buffers Compiler (compilador oficial para arquivos `.proto`)
- Plugins para Go (`protoc-gen-go` e `protoc-gen-go-grpc`)

---

## Como funciona

- Service-a cria um servidor gRPC que envia uma stream contínua de palavras aleatórias.
- Service-b é o cliente que se conecta ao service-a via gRPC e imprime as palavras recebidas.

---

## Como rodar

- Esse projeto é muito simples de executar, bastando apenas gerar os arquivos .proto atraves do comando make all, criar 2 containers (service-a e service-b) e subir os containers:

## Gerar os arquivos .proto

- Para criar os arquivos gerados a partir dos .proto, rode na raiz do projeto:

```bash
make all
```
### Alternativa, rodar o comando direto na raiz:

```bash
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative \
  --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/*.proto
```

---

## Build manual dos containers
- Para simular um fluxo mais real, gere as imagens individualmente:

```bash
docker build -t service-a -f service-a/Dockerfile .
docker build -t service-b -f service-b/Dockerfile .
```

## Subir os containers
- Execute:

```bash
docker compose up
```

## Exemplo de saída no terminal

```bash
service-a  | 2025/05/21 17:56:13 🔌 gRPC server listening on port 50051
service-b  | 2025/05/21 17:56:14 ✅ Connected to service-a
service-b  | 2025/05/21 17:56:14 📡 Receiving stream of random words from service-a:
service-b  | 2025/05/21 17:56:14 📝 Received: docker
service-b  | 2025/05/21 17:56:19 📝 Received: protobuf
service-b  | 2025/05/21 17:56:24 📝 Received: grpc
service-b  | 2025/05/21 17:56:29 📝 Received: golang
service-b  | 2025/05/21 17:56:34 📝 Received: grpc
```