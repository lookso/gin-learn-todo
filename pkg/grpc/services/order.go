package services

import (
	"context"
	"fmt"
	"gin-learn-todo/pkg/grpc/proto"
)

type OrdersService struct {

}

func(o *OrdersService)NewOrder(ctx context.Context,orderRequest *proto.OrderRequest) (*proto.OrderResponse, error)   {
	// 请求参数
	//{
	//	"order_id": 1,
	//	"order_no": "2",
	//	"user_id": 3,
	//	"order_money": 4,
	//	"order_detail": [
	//	{
	//		"order_id": 1,
	//		"prod_id": "123"
	//	},
	//	{
	//		"order_id": 1,
	//		"prod_id": "123"
	//	}]
	//}
	fmt.Println(orderRequest.OrderMain)
	return &proto.OrderResponse{
		Status:"OK",
		Message:"success",
	},nil
}