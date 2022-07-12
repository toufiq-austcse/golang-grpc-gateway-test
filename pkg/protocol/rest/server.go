package rest

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	todoServicePb "github.com/toufiq-austcse/golang-grpc-gateway-test/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func RunServer(ctx context.Context, grpcPort string, httpPort string) error {
	fmt.Println("dfs")
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := todoServicePb.RegisterTodoServiceHandlerFromEndpoint(ctx, gwMux, "localhost:"+grpcPort, opts)
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: gwMux,
	}
	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
