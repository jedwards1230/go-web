# go-web

Basic gRPC server and cli client 

Planning to turn this into a gRPC chatroom either entirely in go or mixed with python

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
