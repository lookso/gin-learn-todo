package scripts

import (
	"context"
	"fmt"
	"gin-learn-todo/pkg/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
)

const (
	ADDRESS = ":9192"
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

func TestOrder(t *testing.T) {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
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
