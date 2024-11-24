# Working wit gRPC

Quickstart:
https://grpc.io/docs/languages/go/quickstart/

Add your protobuff definitions and run to generate protobuf code:
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/message/\*.proto
