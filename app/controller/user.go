/*
@Time : 2019-11-17 10:21 
@Author : Tenlu
@File : user
@Software: GoLand
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InsterUser(c *gin.Context) {

}
func BatchInsertUser(c *gin.Context) {

}

func Del(c *gin.Context) {

}

func FindUser(c *gin.Context) {
	//fmt.Fprintf(os.Stderr, "123")
	c.JSON(http.StatusOK, time.Now().Format("2006-01-02 15:03:04"))
}

func UpdateUser(c *gin.Context) {

}
