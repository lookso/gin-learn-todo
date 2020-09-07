package user

import (
	"encoding/json"
	"fmt"
	"gin-learn-todo/libs/mysql"
	"gin-learn-todo/libs/redis"
	"gin-learn-todo/model"
	"gin-learn-todo/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// @Summary 用户详情接口
// @Description 用户详情
// @Tags 业务API
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /user/info/{id} [get]
// @Security ApiKeyAuth
// @Param id path int true "用户id"
// @Success 200 {object} model.User "用户详情"
func Info(c *gin.Context) {
	var err error
	var user model.User
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
	userInfoKey := fmt.Sprintf(redis.SnsUserInfoKey, uid)
	if userStr, err = redis.Client().Get(userInfoKey).Result(); err != nil && err != redis.ErrNil {
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
		return
	}
	c.JSON(response.Data(user))
}

// @Summary 用户列表
// @Description 用户列表
// @Tags 业务API
// @Accept x-www-form-urlencoded
// @Produce json
// @Router /user/list [get]
// @Security ApiKeyAuth 
// @Success 200 {array} model.User "用户列表"
func List(c *gin.Context) {
	var err error
	var users []model.User
	if err = mysql.Client().Where("status=1").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(response.ParamsError("数据不存在"))
		return
	}
	var count int32
	if err = mysql.Client().Model(&model.User{}).
		Where("status=1").Count(&count).Error; err != nil {
		c.AbortWithStatusJSON(response.ParamsError("获取总数失败"))
		return
	}
	c.JSON(response.DataWithTotal(count, users))
}
