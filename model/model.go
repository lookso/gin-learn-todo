/*
@Time : 2020-01-31 16:01 
@Author : peanut
@File : model
@Software: GoLand
*/
package model

type Model struct {
	Id         uint  `json:"id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}
