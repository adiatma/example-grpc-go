package main

import (
	"example-grpc-go/product-grpc/handler"
	pb "example-grpc-go/protos"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	network = "tcp"
	port    = ":50051"
)

func main() {
	// listen tcp server
	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatalf("listen server error: %v", err)
	}

	// define grpc server
	server := grpc.NewServer()

	// Enable server reflection
	// @see https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(server)

	// register service
  pb.RegisterProductServiceServer(server, handler.Server())

	fmt.Printf("Server running in %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Printf("error %v", err)
	}
}
