# Install `grpc`

Install packages:
1. enable `protoc --go_out` to generate messages code  
`go install google.golang.org/protobuf/cmd/protoc-gen-go`

1. enable `protoc --go-grpc_out` to generate service stub code from `service` block  
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`