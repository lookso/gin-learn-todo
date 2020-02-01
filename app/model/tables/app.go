/*
@Time : 2020-02-01 14:12 
@Author : peanut
@File : app
@Software: GoLand
*/
package tables

import (
	"gin-learn-todo/app/resources/mysql"
	"github.com/jinzhu/gorm"
)

type App struct {
	gorm.Model

	Name    string `json:"name"`
	Status  int    `json:"status"`
	Desc    string `json:"desc"`
	Creator string `json:"creator"`
}

func (a *App) GetAppById(appId string) (*App, error) {
	var app App
	if err := mysql.Client().Where("id=? and status=1", appId).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

