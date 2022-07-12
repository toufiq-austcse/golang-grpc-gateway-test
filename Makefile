create:
	protoc -I ./api/proto \
       --go_out ./pkg/api --go_opt paths=source_relative \
       --go-grpc_out ./pkg/api --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./pkg/api --grpc-gateway_opt paths=source_relative \
       ./api/proto/v1/todo-service.proto