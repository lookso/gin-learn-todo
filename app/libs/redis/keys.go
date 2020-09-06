/*
@Time : 2020-02-22 18:53 
@Author : peanut
@File : keys
@Software: GoLand
*/
package redis

// redis key 命名格式：多个单词之间用冒号(:)连接，碰到由单个单词组成的词组用下划线(_)连接 const

const (
	SnsUserInfoKey = "sns:user_info:id:%d"
)