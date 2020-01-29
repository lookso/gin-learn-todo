package controller

import (
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
	c.JSON(response.Data(data))
}

func List(c *gin.Context)  {
	c.String(http.StatusOK,"ok")
}
