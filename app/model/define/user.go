/*
@Time : 2020-01-31 10:22 
@Author : peanut
@File : post
@Software: GoLand
*/
package define

type User struct {
	// gorm.Model
	Model
	UserName  string `json:"username"`
	Passport  string `json:"passport"`
	Phone     int64  `json:"phone"`
	NickName  string `json:"nick_name"`
	Avatar    string `json:"avatar"`
	Medium    string `json:"medium"`
	Age       int    `json:"age"`
	Sex       int    `json:"sex"`
	Astrology string `json:"astrology"`
	Status    int    `json:"status"`
	Source    string `json:"source"`
}
