package grpc

import (
	"gin-learn-todo/pkg/grpc/proto"
	"gin-learn-todo/pkg/grpc/services"
	"gin-learn-todo/pkg/log"
	"gin-learn-todo/pkg/trace"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
	"net"
)

const (
	PORT = ":9192"
)

func Run() {
	zkTracer, err := trace.NewTrace()
	if err != nil {
		log.Sugar().Fatalf("zipkin trace: %v", err)
	}
	// grpc_middleware 、grpc_recover
	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(zkTracer, otgrpc.LogPayloads())))
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Sugar().Fatalf("failed to listen: %v", err)
	}
	// 注册服务
	proto.RegisterProdServer(rpcServer, new(services.ProdService))
	proto.RegisterOrderServer(rpcServer, new(services.OrdersService))

	rpcServer.Serve(lis)
}
