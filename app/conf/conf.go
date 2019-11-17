/*
@Time : 2019-11-17 17:05 
@Author : Tenlu
@File : conf
@Software: GoLand
*/
package conf

import "gin-learn-todo/app/utils/log"

type Config struct {
	Name   string
	Debug  bool
	Listen string
	Log    log.Config
}


