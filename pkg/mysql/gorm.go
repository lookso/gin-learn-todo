package mysql

import (
	"gin-learn-todo/pkg/log"
	"gin-learn-todo/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	db *gorm.DB
)

// 初始化Mysql数据库连接池
func Init() (err error) {
	db, err = gorm.Open("mysql", setting.Conf.Mysql.Sns.Addr)
	if err != nil {
		log.Sugar().Errorf("init mysql fail %v", err)
		return err
	}
	db.DB().SetMaxIdleConns(setting.Conf.Mysql.Sns.MaxIdleConns)
	db.DB().SetMaxOpenConns(setting.Conf.Mysql.Sns.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(setting.Conf.Mysql.Sns.MaxLifeTime) * time.Second)

	// 如果是本地开发，则开启sql日志
	if setting.Conf.App.Env == "dev" {
		db.LogMode(true)
	}
	log.Sugar().Info("init mysql success")
	return nil
}

func Client() *gorm.DB {
	return db
}

// Close 关闭
func Close() {
	db.Close()
}
