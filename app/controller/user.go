package controller

import (
	"encoding/json"
	"fmt"
	"gin-learn-todo/app/enums/rediskeys"
	"gin-learn-todo/app/model/define"
	"gin-learn-todo/app/resources/mysql"
	"gin-learn-todo/app/resources/redis"
	"gin-learn-todo/pkg/response"
	"github.com/gin-gonic/gin"
	redisLib "github.com/go-redis/redis"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Info(c *gin.Context) {
	var err error
	var user define.User
	id := c.Query("id")
	if id == "" {
		c.AbortWithStatusJSON(response.ParamsError("Id 不能为空"))
		return
	}
	var uid int
	if uid, err = strconv.Atoi(id); err != nil {
		c.AbortWithStatusJSON(response.ParamsError("Id 错误"))
		return
	}
	var userStr string
	userInfoKey := fmt.Sprintf(rediskeys.SnsUserInfoKey, uid)
	if userStr, err = redis.Client().Get(userInfoKey).Result(); err != nil && err != redisLib.Nil {
		log.Fatalf("redis get key(%v) err: %v", userInfoKey, err)
	}
	if userStr != "" {
		err := json.Unmarshal([]byte(userStr), &user)
		if err != nil {
			log.Fatalf("JsonToStruct err:%v", err)
		}
		c.JSON(response.Data(user))
		return
	}
	if err = mysql.Client().Where("id=? and status=1", uid).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(response.ParamsError("数据不存在"))
		return
	}
	var userByte []byte
	userByte, err = json.Marshal(&user)
	if err := redis.Client().Set(userInfoKey, string(userByte), time.Duration(10*60)*time.Second).Err(); err != nil {
		log.Fatalf("redis set key(%v) err: %v", userInfoKey, err)
	}
	c.JSON(response.Data(user))
}

func List(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
