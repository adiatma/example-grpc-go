package handler

import (
	"context"
	pb "example-grpc-go/protos"
)

type (
	ps struct {
		pb.ProductServiceServer
	}
)

func Server() pb.ProductServiceServer {
	ps := new(ps)
	return ps
}

// add service handler
func (ps *ps) GetProduct(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	var products []*pb.Product
	products = append(products, &pb.Product{
		Id:    1,
		Title: "Test",
		Body:  "Test",
	})
	return &pb.ProductResponse{
		Result: products,
	}, nil
}
