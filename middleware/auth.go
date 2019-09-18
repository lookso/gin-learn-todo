package middleware

import "github.com/gin-gonic/gin"

func Login() gin.HandlerFunc {
	// 验证通过，会继续访问下一个中间件
	return func(c *gin.Context) {
		c.Next()
	}
}
