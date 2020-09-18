package middleware

import (
	zk "gin-learn-todo/pkg/trace"
	"github.com/gin-gonic/gin"
)

//添加一个 middleWare, 为每一个请求添加span
func ZipKin() gin.HandlerFunc {
	return func(c *gin.Context) {
		span := zk.ZkTracer.StartSpan(c.FullPath())
		defer span.Finish()

		c.Next()
	}
}
