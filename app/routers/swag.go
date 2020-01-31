/*
@Time : 2020-01-30 21:02 
@Author : peanut
@File : swag
@Software: GoLand
*/
package routers

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
