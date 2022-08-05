# Blueprint ROA Golang gRPC Gateway With Emitter Kafka Streams

## Description

Blueprint Emitter Service

## Documentation

Emitter is one of the 3 parts of Goka's Kafka streams that produce a message to
specified stream topic with given configurations. When Emitter producing
a message, there may be a `Processor` waiting to process the message from stream topic
such as counter, filter or lookup etc.

## Features

* **Producer Message**
* **Support gRPC and RESTful**

## Dependencies

- Goka
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

## Prerequisites

- Docker
- Kafka
- Zookeeper
- Postgres
- Protoc

## Run locally

1. Install protoc gen gRPC gateway

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

2. Install protoc gen OpenAPI (Swagger)

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

3. Compile `.proto` file

```shell
make pbgen
```

3. Start server

```shell
make start
```

### Recommendation

Before push, should run linter always for code quality

1. Install linter (golangci-lint)

```shell
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

2. Run linter

```shell
golangci-lint run
```

## Connections

- GRPC: localhost:3000
- HTTP: localhost:3001
- Swagger: localhost:3001/swagger-ui

## Resources

- [Apache Kafka](https://kafka.apache.org)
- [Goka's GoDoc](https://godoc.org/github.com/lovoo/goka)
- [Golangci-lint](https://golangci-lint.run/)

## Happy Coding...