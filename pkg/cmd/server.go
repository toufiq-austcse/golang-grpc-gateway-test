package cmd

import (
	"context"
	"fmt"
	"github.com/toufiq-austcse/golang-grpc-gateway-test/pkg/protocol/grpc"
	"github.com/toufiq-austcse/golang-grpc-gateway-test/pkg/protocol/rest"
	TodoServiceServer "github.com/toufiq-austcse/golang-grpc-gateway-test/pkg/service/v1"
)

func RunServer() error {
	ctx := context.Background()
	grpcPort := "8080"
	httpPort := "9090"
	go func() {
		//run http gatway
		_ = rest.RunServer(ctx, grpcPort, httpPort)
	}()
	if err := grpc.RunServer(TodoServiceServer.NewTodoServiceServer(), grpcPort); err != nil {
		return err
	}
	fmt.Println("Listening on grpcPort " + grpcPort)

	return nil
}
