package controller

import (
	"github.com/gin-gonic/gin"

	"code.itech8.com/openapi/sns-manager-api/app/helpers/response"
)

type Post struct{}

func (p *Post) List(c *gin.Context) {
	data := map[string]interface{}{"name": "jack"}
	response.Success(data, c)
}

func (p *Post) Detail(c *gin.Context) {
	data := map[string]interface{}{"name": "jack"}
	response.Success(data, c)
}
// 帖子审核
func (p *Post) Audit(c *gin.Context) {
	data := map[string]interface{}{"name": "jack"}
	response.Success(data, c)
}
