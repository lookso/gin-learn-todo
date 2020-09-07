/*
@Time : 2020-02-22 18:21 
@Author : peanut
@File : alarm
@Software: GoLand
*/
package resource

import (
	"time"
)

type Alarm struct {
	MainTitle   string    `json:"main_title"`    // 主标题
	SubTitle    string    `json:"sub_title"`     // 副标题
	CoverImg    string    `json:"cover_img"`     // cover图片
	UnlockType  string    `json:"unlock_type"`   // 解锁类型 '解锁类型:免费:free 快乐币解锁:coin 快乐能量解锁:energy 视频广告解锁 ad'
	PayMoney    int64     `json:"pay_money"`     // 付费金额
	Tag         string    `json:"tag"`           // new 最新,hot 最热
	SetPlayNum  int       `json:"set_play_num"`  // 后台设置的播放量
	RealPlayNum int       `json:"real_play_num"` // 用户真实播放量
	Duration    int       `json:"duration"`      // 闹铃音频时长
	AlarmUrl    string    `json:"alarm_url"`     // 闹铃地址
	StartTime   time.Time `json:"start_time"`    // 下发时间
	AlarmType   string    `json:"alarm_type"`    // 音频:audio 视频:video
}

