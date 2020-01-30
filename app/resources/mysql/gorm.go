package mysql

import (
	"fmt"
	"gin-learn-todo/app/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

var (
	db *gorm.DB
)

// 初始化Mysql数据库连接池
func Init() (db *gorm.DB, err error) {
	conn, err := gorm.Open("mysql", setting.Conf.Mysql.Sns.Addr)
	if err != nil {
		errStr := fmt.Sprintf("mysql: sql.Open() error,%v", err)
		log.Fatal(errStr)
		return nil, err
	}
	conn.DB().SetMaxIdleConns(setting.Conf.Mysql.Sns.MaxIdleConns)
	conn.DB().SetMaxOpenConns(setting.Conf.Mysql.Sns.MaxOpenConns)
	conn.DB().SetConnMaxLifetime(time.Duration(setting.Conf.Mysql.Sns.MaxLifeTime) * time.Second)

	// 如果是本地开发，则开启sql日志
	if os.Getenv("ENV") == "dev" {
		db.LogMode(true)
	}

	return conn, err
}

// Close 关闭
func Close() {
	db.Close()
}
