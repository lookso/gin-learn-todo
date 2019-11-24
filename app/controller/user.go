package controller

import (
	"github.com/gin-gonic/gin"

	"code.itech8.com/openapi/sns-manager-api/app/helpers/response"
)

type User struct{}

func (u *User) Register(c *gin.Context) {
	data := map[string]interface{}{"name": "jack"}
	response.Success(data, c)
}

func (u *User) Detail(c *gin.Context) {
	data := map[string]interface{}{"name": "jack"}
	response.Success(data, c)
}

func (u *User) Edit(c *gin.Context) {
	data := map[string]interface{}{"name": "jack"}
	response.Success(data, c)
}


