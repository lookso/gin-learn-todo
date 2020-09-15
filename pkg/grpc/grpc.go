package grpc

import (
	"gin-learn-todo/pkg/grpc/proto"
	"gin-learn-todo/pkg/grpc/services"
	"gin-learn-todo/pkg/log"
	"google.golang.org/grpc"
	"net"
)

const (
	PORT = ":9192"
)

func Run() {

	rpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Sugar().Fatalf("failed to listen: %v", err)
	}
	proto.RegisterProdServer(rpcServer, new(services.ProdService))
	proto.RegisterOrderServer(rpcServer, new(services.OrdersService))

	rpcServer.Serve(lis)
}
