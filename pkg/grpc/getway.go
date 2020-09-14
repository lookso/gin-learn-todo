package grpc

import (
	"flag"
	gw "gin-learn-todo/pkg/grpc/proto"
	"gin-learn-todo/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9192", "endpoint of Gateway")
)

func Gateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterProductHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	log.Sugar().Info("grpc gateway init success")
	return http.ListenAndServe(":8080", mux)
}
