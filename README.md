# Spacecore Registry

Implementation of the Spacecore registry using IPFS and libp2p. The Spacecore registry is a decentralized plugin registry for the Vistara network.

### Features

- Register plugins with the Spacecore registry, discover plugins, and get plugin details.
- Store plugin metadata in IPFS and plugin binaries in IPFS UnixFS.
- Use libp2p for peer discovery and communication.
- Use gRPC for the service definition.
- Use Protocol Buffers for the message format.

### Project Structure

```
.
├── README.md
├── cmd
├── go.mod
├── go.sum
├── internal
│   ├── grpc_server.go
│   ├── ipfs
│   │   └── ipfs.go
│   ├── p2p
│   │   └── p2p.go
│   └── plugin_manager.go
├── main.go
├── oldgomod
└── pb
    ├── spacecore.pb.go
    ├── spacecore.proto
    └── spacecore_grpc.pb.go
```

### Define the gRPC service and messages:

```proto
syntax = "proto3";

package pb;

option go_package = "./pb";

service PluginRegistry {
    rpc RegisterPlugin (RegisterPluginRequest) returns (RegisterPluginResponse);
    rpc DiscoverPlugins (DiscoverPluginsRequest) returns (DiscoverPluginsResponse);
    rpc GetPlugin (GetPluginRequest) returns (GetPluginResponse);
}
message Plugin {
    string name = 1;
    string version = 2;
    string cid = 3;
    string path = 4;
}

message RegisterPluginRequest {
    string name = 1;
    string version = 2;
    string plugin = 3;
}

message RegisterPluginResponse {
    string message = 1;
    string cid = 2;
}

message DiscoverPluginsRequest {
    optional string name = 1;
    optional string cid = 2;
}

message DiscoverPluginsResponse {
    repeated Plugin plugins = 1;
}

message GetPluginRequest {
    string name = 1;
    string version = 2;
}

message GetPluginResponse {
    Plugin plugin = 1;
}
```

Let's use buf to generate proto files.

Generate the Go code from the proto file:

```sh
protoc --go_out=. --go-grpc_out=. pb/spacecore.proto
```

If using `buf`
```sh
buf generate
```

### Running the Service

1. **Start IPFS**:
   ```sh
   ipfs daemon
   ```

2. **Run the gRPC Server**:
   ```sh
   go run cmd/main.go
   ```

### Testing the gRPC Service

You can test the gRPC service using a gRPC client like `grpcurl` or by writing a client in Go, Python, etc.

```sh
grpcurl -plaintext -d '{"name": "example", "description": "example description", "container_image": "example/image:latest"}' localhost:50051 pb.SpacecoreRegistry/RegisterSpacecore
```
