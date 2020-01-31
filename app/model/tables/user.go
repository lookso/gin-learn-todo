package tables

import (
	"gin-learn-todo/app/model/define"
)

type UserRes struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Desc      string `json:"desc"`
	Creator   string `json:"creator"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var user define.User

//func (u *user) ToUserJsonArray(users []user) []UserRes {
//	var userRes = make([]UserRes, len(user))
//	for k, v := range app {
//		userRes[k] = AppRes{
//			Id:        v.Id,
//			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
//			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
//		}
//	}
//	return appRes
//}
