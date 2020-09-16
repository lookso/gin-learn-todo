package grpc

import (
	"flag"
	"gin-learn-todo/pkg/grpc/proto"
	"gin-learn-todo/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9192", "endpoint of Gateway")
)
// 启用grpc gateway
func Gateway() {
	gwMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterProdHandlerFromEndpoint(context.Background(),
		gwMux, *echoEndpoint, opt)
	if err != nil {
		log.Sugar().Fatalf("grpc gateway register order-server err %v", err)
	}
	err = proto.RegisterOrderHandlerFromEndpoint(context.Background(),
		gwMux, *echoEndpoint, opt)
	if err != nil {
		log.Sugar().Fatalf("grpc gateway register product-server err %v", err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwMux,
	}
	httpServer.ListenAndServe()
}
