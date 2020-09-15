package services

import (
	"context"
	"gin-learn-todo/pkg/grpc/proto"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, request *proto.ProdRequest) (*proto.ProdResponse, error) {

	return &proto.ProdResponse{ProdStock: 20}, nil
}

func (p *ProdService) GetProdStocks(ctx context.Context, size *proto.QuerySize) (*proto.ProdResponseList, error) {
	var Prodres []*proto.ProdResponse
	Prodres = make([]*proto.ProdResponse, 0, 3)
	Prodres = append(Prodres, &proto.ProdResponse{ProdStock: 28})
	Prodres = append(Prodres, &proto.ProdResponse{ProdStock: 29})
	Prodres = append(Prodres, &proto.ProdResponse{ProdStock: 30})
	return &proto.ProdResponseList{
		Prodres: Prodres,
	}, nil
}

func (p *ProdService) GetProdInfo(ctx context.Context, in *proto.ProdRequest) (*proto.ProdModel, error) {

	ret := proto.ProdModel{
		ProdId:    101,
		ProdName:  "测试商品",
		ProdPrice: 20.5,
	}
	return &ret, nil
}
