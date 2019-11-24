package mysql

import (
	"code.itech8.com/openapi/sns-manager-api/app/config"
	"code.itech8.com/openapi/sns-manager-api/app/helpers/xlog"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/url"
	"time"
)

var (
	//DBJokeSns
	DbSns *gorm.DB
)

func init() {
	dbConfList := config.Config().Db
	DbSns = initDbConnByConf(genDbDsn(dbConfList.Sns), dbConfList.Sns)
}

// 初始化Mysql数据库连接池
func initDbConnByConf(dsn string, dbConf *config.DbConf) *gorm.DB {
	xlog.Bare("dsn", dsn)
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		errStr := fmt.Sprintf("mysql: sql.Open() error, %v", err)
		xlog.Error(errStr)
	}
	conn.DB().SetMaxOpenConns(dbConf.MaxOpenConns)
	conn.DB().SetMaxIdleConns(dbConf.MaxIdleConns)
	conn.DB().SetConnMaxLifetime(time.Duration(dbConf.MaxLifeTime) * time.Second)
	xlog.Bare("gorm专用链接 - 数据库连接池初始化完成：", dbConf.DbName)

	return conn
}

// DSN 说明： https://github.com/go-sql-driver/mysql#system-variables
// 关于mysql连接的时区 和 表里 date/datetime 类型的字段映射为time.Time时的时区：
// 	loc 只对 time.Time 类型的表字段转换有效，不会修改 mysql 连接的时区设置；
// 	time_zone 设置mysql连接的默认时区

func genDbDsn(dbConf *config.DbConf) (dsn string) {
	xlog.Bare("gorm专用连接 - 当前数据库名：", dbConf.DbName)

	loc := "Local"
	//timeZone := url.QueryEscape(Config().BmsServe.Timezone)
	timeZone := url.QueryEscape("'+0:00'")
	//loc = timeZone
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s&time_zone=%s&interpolateParams=true",
		dbConf.User, dbConf.Pwd, dbConf.Host, dbConf.Port, dbConf.DbName, dbConf.Charset, loc, timeZone)
	return
}
