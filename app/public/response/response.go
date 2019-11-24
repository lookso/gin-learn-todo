/*
@Time : 2019/3/7 5:00 PM
@Author : Tenlu
@File : response
@Software: GoLand
*/
package response

import (
	"code.itech8.com/openapi/sns-manager-api/app/helpers/xlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func Success(data interface{}, c *gin.Context) {
	errMsg := CodeMsg(RequertSuccess)
	c.JSON(http.StatusOK, &Error{RequertSuccess, errMsg, data})
	c.Abort()
}

func Fail(errorCode int, c *gin.Context) {
	errMsg := CodeMsg(errorCode)
	c.JSON(http.StatusBadRequest, &Error{errorCode, errMsg, map[string]interface{}{}})
	//c.Abort()
	return
}


func CodeMsg(code int) string {
	if _, ok := errorMsg[code]; !ok {
		xlog.Error("no match error, error code:", code)
		return ""
	}
	return errorMsg[code]
}