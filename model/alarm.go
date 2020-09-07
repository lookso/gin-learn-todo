/*
@Time : 2020-02-22 18:22 
@Author : peanut
@File : alarm
@Software: GoLand
*/
package model

import (
	"github.com/jinzhu/gorm"
)

type Alarm struct {
	gorm.Model
	MainTitle   string    `json:"main_title"`    // 主标题
	Status      int8      `json:"status"`        // 1 上线 2下线
	IsTop       int       `json:"is_top"`        // 是否置顶0否1是
}

// Callbacks

// 此方法不需要显示的调用，会在程序调用 Create() 方法前自动调用
func (alarm *Alarm) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Status", 1)
	return nil
}
func (alarm *Alarm) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn("Status", 1)
	return nil
}
func (alarm *Alarm) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Status", 1)
	return nil
}

func (alarm *Alarm) AfterFind() {
	alarm.IsTop = 1
}
