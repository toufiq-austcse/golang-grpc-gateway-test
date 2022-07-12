package grpc

import (
	"fmt"
	todoServicePb "github.com/toufiq-austcse/golang-grpc-gateway-test/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunServer(srv todoServicePb.TodoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	fmt.Println(listen, err)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	server := grpc.NewServer()
	todoServicePb.RegisterTodoServiceServer(server, srv)
	log.Println("Serving gRPC on 0.0.0.0:" + port)
	return server.Serve(listen)
}
