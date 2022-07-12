package v1

import (
	"context"
	todoServicePb "github.com/toufiq-austcse/golang-grpc-gateway-test/pkg/api/v1"
)

type todoServiceServer struct {
	todoServicePb.UnimplementedTodoServiceServer
}

func NewTodoServiceServer() *todoServiceServer {
	return &todoServiceServer{}
}
func (s *todoServiceServer) Create(ctx context.Context, in *todoServicePb.CreateTodoRequest) (*todoServicePb.CreateTodoResponse, error) {
	return &todoServicePb.CreateTodoResponse{Id: 2}, nil
}
