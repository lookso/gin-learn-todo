package controller

import (
	"fmt"
	"gin-learn-todo/pkg/response"
	"github.com/gin-gonic/gin"
)

func GetRawData(c *gin.Context) {
	rawData, err := c.GetRawData() // 获取原始数据流
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(response.Data(rawData))
}
