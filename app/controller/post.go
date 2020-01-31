package controller

import (
	"fmt"
	"gin-learn-todo/app/resources/redis"
	"gin-learn-todo/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Info(c *gin.Context) {
	data := struct {
		Name string `json:"name"`
	}{
		Name: "hello gin",
	}
	if err := redis.Client().Set("timestamp", time.Now().Format("2006-01-02 15:04:05"), 0).Err();err!=nil{
		fmt.Println(err)
	}
	timestamp,_:=redis.Client().Get("timestamp").Result()
	fmt.Println(timestamp)
	

	c.JSON(response.Data(data))
}

func List(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
