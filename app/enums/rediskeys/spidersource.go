/*
@Time : 2019-09-21 20:52 
@Author : Tenlu
@File : spidersource
@Software: GoLand
*/
package rediskeys

// redis key 命名格式：多个单词之间用冒号(:)连接，碰到由单个单词组成的词组用下划线(_)连接 const

const (
	// 爬虫数据源站点介绍详情
	SpiderSourceDetailKey = "sns:spidersource:detail:id:%d"
	SpiderSourceListKey = "sns:spidersource:list:id:%d"
)
