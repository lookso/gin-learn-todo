package grpc

import (
	"context"
	"fmt"
	"gin-learn-todo/pkg/grpc/proto"
	"gin-learn-todo/pkg/log"
	"google.golang.org/grpc"
	"net"
)

const (
	PORT = ":9192"
)

type server struct{}

func (s *server) ProductInfo(ctx context.Context, in *proto.ProductInfoRequest) (*proto.ProductInfoResponse, error) {
	fmt.Println(in.Id)
	return &proto.ProductInfoResponse{Id: 1, Name: "peanut", Price: 32}, nil
}
func (s *server) CreateProduct(ctx context.Context, in *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	fmt.Println(in.Name, in.Price)
	return &proto.CreateProductResponse{Status: "success"}, nil
}

func Run() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Sugar().Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterProductServer(s, &server{})
	log.Sugar().Info("grpc start success")
	s.Serve(lis)
}
