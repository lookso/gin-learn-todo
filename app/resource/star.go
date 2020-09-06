/*
@Time : 2020-02-22 21:43 
@Author : peanut
@File : star
@Software: GoLand
*/
package resource

type Star struct {
	Name     string  `json:"name"`      // 明星合集
	Sort     int     `json:"sort"`      // 排序
	CoverImg string  `json:"cover_img"` // 明星cover图片
	Version  string  `json:"version"`   // 版本号
	Status   int8    `json:"status"`    // 1 上线 2下线
	IsIndex  int     `json:"is_index"`  // 是否在首页 0 否 1 是
	Alarm    []Alarm `json:"alarms"`    // 一对多的关联关系(一个明星有多个闹铃)
}
