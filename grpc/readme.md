#### install grpc
```go
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```


#### generating proto
SHOULD NOT USE `--go-out=plugin` DONOT MISS THE `S`
```go
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

### Steps
- First create the proto required
- Generate the go file using using protoc
- Create server
    - create listener & register new service
    - implement the server interface
- Create client
    - dial up connection & create new service client
    - implement the client interface

### steps in creating a server LSRS
- Listener -> create a listener
- Server -> Create a server
- Register -> Register the service to the server
- Serve -> Server on the listener port

### steps in creating a client
- Dial -> grpc dial connection
- Client -> Create new service client
- Request -> create request object
- Response -> get resonse from ServiceClient interface

### JSON to gRPC
https://github.com/grpc-ecosystem/grpc-gateway