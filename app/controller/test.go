package controller

import (
	"gin-learn-todo/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func DataInfo(c *gin.Context) {
	time.Sleep(5 * time.Second)
	data := map[string]string{"name": "jack", "tag": "<br>", "北京时间": time.Now().Format("2006-01-02 15:03:04")}

	utils.Success(data, c)
}
func InsertData(c *gin.Context) {
	data := map[string]string{"北京时间": time.Now().Format("2006-01-02 15:04:05")}
	utils.Success(data, c)
}
func task() string {
	// 用 time.Sleep() 模拟一个长任务。
	time.Sleep(5 * time.Second)
	return time.Now().Format("2006-01-02 15:04:05")
}
func GetPing(c *gin.Context) {
	// 创建在 goroutine 中使用的副本
	cCp := c.Copy()
	go func() {
		NowTime := task()
		// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
		log.Println("Done! in path " + cCp.Request.URL.Path + "->" + NowTime)
	}()

	var params = make(map[string]string)
	// c.Query 查询字符串参数

	if name := c.Query("name"); name != "" {
		params["name"] = name
	}
	if age := c.Query("age"); age != "" {
		params["age"] = age
	}
	data := map[string]interface{}{
		"lang":   "GO语言",
		"tag":    "<br>",
		"params": params,
	}
	c.AsciiJSON(http.StatusOK, data)
}
func GetJsonp(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	// callback 是 x
	// 将输出：x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}

func GetHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "template/index.tmpl", gin.H{
		"title": "Go Main Website",
		"body":  "This is Go Body",
	})
}
