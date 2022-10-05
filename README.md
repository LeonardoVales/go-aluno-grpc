# Exemplo de implementação de gRPC

1. Suba o container rodando:
```
docker-compose up -d
```

2. Acesse o container:
```
docker-compose exec app sh
```

3. Para realizar um teste, dentro do container em uma aba do terminal execute o server, e em outra aba execute o client:
```
go run cmd/server/server.go
go run cmd/client/client.go
```

4. Caso queira gerar novamente as stubs:
```
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.

