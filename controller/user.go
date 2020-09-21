package controller

import (
	"encoding/json"
	"fmt"
	"gin-learn-todo/model"
	"gin-learn-todo/pkg/jwt"
	"gin-learn-todo/pkg/log"
	"gin-learn-todo/pkg/mysql"
	"gin-learn-todo/pkg/redis"
	"gin-learn-todo/pkg/response"
	"gin-learn-todo/pkg/utils"
	"gin-learn-todo/resource"
	"github.com/gin-gonic/gin"
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
	log.Sugar().Errorf("zap log test")
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
		log.Sugar().Errorf("redis get key(%v) err: %v", userInfoKey, err)
	}
	if userStr != "" {
		err := json.Unmarshal([]byte(userStr), &user)
		if err != nil {
			log.Sugar().Errorf("JsonToStruct err:%v", err)
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
		log.Sugar().Errorf("redis set key(%v) err: %v", userInfoKey, err)
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

var (
	uuid     string
	username = "peanut"
	password = "123456"
)

func Login(c *gin.Context) {
	var urq resource.LoginRequest
	if err := c.ShouldBind(&urq); err != nil {
		c.JSON(response.BadRequest("参数错误"))
		return
	}
	uuid = utils.GetUuid()
	tokenString, err := jwt.GenerateToken(uuid, urq.UserName, urq.Password)
	if err != nil {
		c.JSON(response.BadRequest("鉴权失败"))
		return
	}
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTg4NjU0NjEsImlhdCI6MTU5ODg2MTg2MX0.qbsLiy9Z_k3J61kuHyxBZYTCY_ZuD2rVmG6zUsY4FBI
	c.JSON(response.Data(tokenString))
}

func UserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.MyClaims)
	if claims.UserName == "" || claims.Uid == "" {
		c.JSON(response.BadRequest(""))
		return
	}
	c.JSON(response.Data(resource.UserInfoResponse{
		Id:   claims.Uid,
		Name: claims.UserName,
		Sex:  1,
		Age:  28,
	}))
}

// 更新token
func Refresh(c *gin.Context) {
	c.JSON(response.Data(nil))
}
