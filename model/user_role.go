/*
@Time : 2020-02-01 14:07 
@Author : peanut
@File : user_role
@Software: GoLand
*/
package model

import (
	"gin-learn-todo/pkg/mysql"
	"github.com/jinzhu/gorm"
)

type UserRole struct {
	gorm.Model
	Email      string `json:"email"`
	RoleId     int    `json:"role_id"`
	AppId      int    `json:"app_id"`
	Status     int    `json:"status"`
	Creator    string `json:"creator"`
	Desc       string `json:"desc"`
}

func (ur *UserRole) GetRolesByEmail(appId string, email string) ([]UserRole, error) {
	var userRoleList []UserRole
	err := mysql.Client().Where("app_id=? and status=1 and email=?", appId, email).Find(&userRoleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return userRoleList, nil
}

func Verifier(appId, email, path, method string) (bool, error) {

	var userRole UserRole
	uRoles, err := userRole.GetRolesByEmail(appId, email)

	if err != nil {
		return false, err
	}
	var isPass = false
	var roleIds = make([]int, len(uRoles))
	for k, ur := range uRoles {
		roleIds[k] = ur.RoleId
	}
	type Result struct {
		AppId  int
		Method string
		Path   string
		RoleId int
	}
	var result []Result
	err = mysql.Client().Raw("select r.`app_id`,`method`,`path`,role_id from resources_roles rs left join resources  r on r.id=rs.resource_id "+
		"where role_id in(?) and r.app_id=? and r.status=1 and rs.app_id=? and rs.status=1 and path=? and method=?", roleIds, appId, appId, path, method).Scan(&result).Error
	if err != nil {
		return false, err
	}
	if len(result) > 0 {
		return true, nil
	}
	return isPass, nil
}
