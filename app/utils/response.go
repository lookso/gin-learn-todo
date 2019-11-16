package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code   int         `json:"code"` // 一定要首字母大写,可导出
	ErrMsg string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, &Error{Code: 0, ErrMsg: "success", Data: data})
	c.Abort()
}
