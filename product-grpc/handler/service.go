package handler

import (
	"context"
	"errors"
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

var products []*pb.Product

func (ps *ps) GetProducts(ctx context.Context, in *pb.Empty) (*pb.ProductsResponse, error) {
	// insert new produt
	products = append(products, &pb.Product{
		Id:    1,
		Title: "Lorem ipsume sit amet dolor",
		Body:  "Lorem ipsume sit amet dolor",
	})
	// return the products
	return &pb.ProductsResponse{
		Result: products,
	}, nil
}

func (ps *ps) GetProduct(ctx context.Context, in *pb.ProductId) (*pb.ProductResponse, error) {
	// loop each product and find and match product id
	for _, product := range products {
		if in.GetId() == product.Id {
			return &pb.ProductResponse{
				Result: product,
			}, nil
		}
	}
	// return error if product == nil
	return nil, errors.New("Error")
}

func (ps *ps) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductResponse, error) {
	// get new product
	newProduct := &pb.Product{
		Id:    in.GetId(),
		Title: in.GetTitle(),
		Body:  in.GetBody(),
	}
	// insert new product to products variable
	products = append(products, newProduct)
	// return the product
	return &pb.ProductResponse{
		Result: newProduct,
	}, nil
}
