# go-web

Basic gRPC server and cli client 

### Start Server (default port 8090)

```
go run server.go
```

### Start Client (default port 8090)

```
go run client.go
```

### Rebuild protobuffer
```
protoc --go_out=plugins=grpc:. chat.proto
```
