/*
@Time : 2019-11-24 12:13 
@Author : Tenlu
@File : mysql
@Software: GoLand
*/
package mysql

import "database/sql"

var (
	Client *sql.DB
)

func init() {
	
	Client, _ = sql.Open("mysql", "root:passwd@itech8$@tcp(127.0.0.1:3306)/test?charset=utf8")
	Client.SetMaxOpenConns(2000)
	Client.SetMaxIdleConns(1000)
	Client.Ping()
}

