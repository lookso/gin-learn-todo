package main

import (
	"gin-learn/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code   int         `json:"code"` // 一定要首字母大写,可导出
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, &Error{Code: 0, ErrMsg: "success", Data: data})
	c.Abort()
}
func UserList(c *gin.Context) {
	data := map[string]string{"name": "jack"}
	Success(data, c)
}
func main() {
	router := gin.New()
	router.Use(middleware.Login())

	userGroup := router.Group("/user")
	userGroup.GET("/list", UserList)
	router.Run(":8989")
}
