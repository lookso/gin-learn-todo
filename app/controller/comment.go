/*
@Time : 2019-09-21 21:41 
@Author : Tenlu
@File : comment
@Software: GoLand
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Comment struct {
}

func (m *Comment) List(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"nb": "the is nb"})
}
