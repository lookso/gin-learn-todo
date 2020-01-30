package controller

import (
	"fmt"
	"gin-learn-todo/app/resources/redis"
	"gin-learn-todo/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(c *gin.Context) {
	data := struct {
		Name string `json:"name"`
	}{
		Name: "hello gin",
	}

	//_, err := redis.Client.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"), 0).Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println("redis", redis.Client)

	c.JSON(response.Data(data))
}

func List(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
