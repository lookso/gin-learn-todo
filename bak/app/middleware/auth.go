package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(loginType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证通过，会继续访问下一个中间件
		if loginType != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "登录类型错误",
				"data": "",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
