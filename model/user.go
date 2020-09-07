package model

type User struct {
	Model
	UserName   string `json:"username"`    //用户姓名
	Passport   string `json:"passport"`    //用户密码
	Phone      int    `json:"phone"`       //手机号码
	Nickname   string `json:"nickname"`    //昵称
	Avatar     string `json:"avatar"`      //小头像
	Medium     string `json:"medium"`      //大头像
	Age        int    `json:"age"`         // 年龄
	Sex        int    `json:"sex"`         // 性别:1男2女
	Astrology  string `json:"astrology"`   // 星座
	Status     int    `json:"status"`      // 1有效 2无效
	Source     string `json:"source"`      // 来源 0.真实用户 1.糗事百科 2.段子手
	CreateTime int    `json:"create_time"` // 创建时间
	UpdateTime int    `json:"update_time"` // 更新时间
}
