## Installation

* Protocol Buffer Compiler Installation `protoc` 
  * The protocol buffer compiler, `protoc`, is used to compile .proto files, which contain service and message definitions.

```bash
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```



* Go Plugins
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Update your PATH so that the protoc compiler can find the plugins:

```$ export PATH="$PATH:$(go env GOPATH)/bin"```


## Run the example
From the examples/helloworld directory:

1. Compile and execute the server code:

```
$ go run greeter_server/main.go
```


2. From a different terminal, compile and execute the client code to see the client output:

```
$ go run greeter_client/main.go
Greeting: Hello world
```

Congratulations! You’ve just run a client-server application with gRPC.


## Update the gRPC service
1. Open `helloworld/helloworld.proto` and add a new `SayHelloAgain()` method, with the same request and response types:

## Regenerate gRPC code
Before you can use the new service method, you need to recompile the updated .proto file.

While still in the `helloworld` directory, run the following command:
```bash 
 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```


## Update and run the application
You have regenerated server and client code, but you still need to implement and call the new method in the human-written parts of the example application.

### Update the server
Open greeter_server/main.go and add the following function to it:

```
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}
```
### Update the client
Open greeter_client/main.go to add the following code to the end of the main() function body:

```
r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
if err != nil {
log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetMessage())
```
##