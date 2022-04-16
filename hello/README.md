# GRPC Hello Demo
## 1. ```.proto``` file
### 1.1 Specific package name in ```.proto``` file
```proto
option go_package = "/helloworld";

package helloworld;
```
### 1.2 Use the protoc command below
```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```