package middleware

import (
	"gin-learn-todo/libs/jwt"
	"gin-learn-todo/pkg/response"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(response.Forbidden(""))
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(response.UnAuthorized("Token is not valid", nil))
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}