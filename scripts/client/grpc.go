package client

import (
	"context"
	"fmt"
	"gin-learn-todo/pkg/grpc/proto"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	GRpcAddress = ":9192"
)

var details = []*proto.OrderDetail{{
	ProdId:  1,
	OrderNo: "10086",
	ProdNum: 10,
}}
var om = &proto.OrderMain{
	OrderId:      1,
	OrderNo:      "10086",
	UserId:       2,
	OrderMoney:   3,
	OrderDetails: details,
}

func GRpcClient() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//tracer := opentracing.GlobalTracer() // 全局tracer
	tracer, err := NewTrace()
	fmt.Println(tracer)
	// 连接到 gRPC 服务器
	conn, err := grpc.DialContext(ctx, GRpcAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
	)
	if err != nil {
		log.Fatalf("connect error: %v", err)
		panic(err)
	}
	defer conn.Close()
	// 初始化 gRPC 客户端
	client := proto.NewOrderClient(conn)
	ordCtx, cel := context.WithTimeout(context.Background(), time.Second*5)
	defer cel()

	resp, err := client.NewOrder(ordCtx, &proto.OrderRequest{
		OrderMain: om,
	})

	if err != nil {
		log.Fatalf("grpc neworder error: %v", err)
		fmt.Println(err)
	}
	fmt.Println(resp)
}
