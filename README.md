# gRPC demo with bidirectional streaming

Before getting started, please ensure that you have the following dependencies
installed:

## The protobuf compiler

### OpenBSD
```sh
pkg_add protobuf
```

### Ubuntu Linux
```bash
sudo apt install -y protobuf-compiler
```

### Mac OS
```bash
brew install protobuf
```

## The protobuf compiler plugins
These are required for generating the go-specific code.

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Running the demo

Before compilation, generate the protobuf specific code via:

```sh
make proto
```

Now, run the following commands:

### Server
```sh
make server
```

### Client
```sh
make client
```
