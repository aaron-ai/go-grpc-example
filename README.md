# Go gRPC Example

## Generate Code by IDL

step into the idl directory, then execute:

```shell
protoc --go_out=../ --go_opt=paths=import --go-grpc_out=../ --go-grpc_opt=paths=import ./apache/rocketmq/v2/definition.proto ./apache/rocketmq/v2/admin.proto ./apache/rocketmq/v2/service.proto ./google/protobuf/timestamp.proto ./google/protobuf/duration.proto
```