package client

import (
	"context"
	"fmt"
	"gin-learn-todo/pkg/grpc/proto"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"time"

	//"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"log"
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
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()

	//tracer := opentracing.GlobalTracer() // 全局tracer
	tracer, err := NewTracer()
	if err != nil {
		log.Println(err)
	}
	// 连接到 gRPC 服务器
	conn, err := grpc.DialContext(ctx,GRpcAddress, grpc.WithInsecure(),grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))
	if err != nil {
		log.Fatalf("connect error: %v", err)
		panic(err)
	}
	defer conn.Close()

	// 初始化 gRPC 客户端
	client := proto.NewOrderClient(conn)
	resp, err := client.NewOrder(context.Background(), &proto.OrderRequest{
		OrderMain: om,
	})
	if err != nil {
		log.Fatalf("Getall error: %v", err)
	}
	fmt.Println(resp)
}
