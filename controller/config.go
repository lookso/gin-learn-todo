package controller

import (
	"context"
	"fmt"
	"gin-learn-todo/pkg/etcd"
	"gin-learn-todo/pkg/response"
	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func SetConfig(c *gin.Context) {
	if err := etcd.MustInit(); err != nil {
		c.AbortWithStatusJSON(response.ServerError(""))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//withPrevKV()是为了获取操作前已经有的key-value
	putResp, err := etcd.Cli.Put(ctx, "/config/golang", "golang.com", clientv3.WithPrevKV())
	if err != nil {
		c.AbortWithStatusJSON(response.ServerError(""))
		return
	}
	if putResp.PrevKv != nil {
		fmt.Println("old:", string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
		_, err := etcd.Cli.Put(ctx, "/config/etcd", "etcd.io")
		if err != nil {
			c.AbortWithStatusJSON(response.ServerError(""))
			return
		}
	}
	return
}

type GetConfReq struct {
	Name string `json:"name"`
}

func GetConfig(c *gin.Context) {
	getConfReq := new(GetConfReq)
	if err := c.Bind(getConfReq); err != nil {
		c.AbortWithStatusJSON(response.ServerError(""))
		return
	}
	if getConfReq.Name==""{
		c.AbortWithStatusJSON(response.ParamsError(""))
		return
	}

	if err := etcd.MustInit(); err != nil {
		c.AbortWithStatusJSON(response.ServerError(""))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := etcd.Cli.Get(ctx, "/config", clientv3.WithPrefix())
	if err != nil {
		c.AbortWithStatusJSON(response.ServerError(""))
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	return
}
