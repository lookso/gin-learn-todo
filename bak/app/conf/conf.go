/*
@Time : 2019-11-17 17:05 
@Author : Tenlu
@File : conf
@Software: GoLand
*/
package conf

import (
	"flag"
	zapLog "gin-learn-todo/app/utils/log"
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
	"sync"
)

var conf *tomlConfig
var tomlConfFile string
var tomlConfLoadOnce sync.Once

func init() {
	var confPointer = flag.String("conf", "conf ../config/app_dev.toml", "-conf ../config/app_dev.toml")
	flag.Parse()
	configFile, err := filepath.Abs(*confPointer)
	if err != nil {
		log.Printf("parse config err(%v)", err)
	}
	tomlConfFile = configFile
}

func Config() *tomlConfig {
	tomlConfLoadOnce.Do(func() {
		if _, err := toml.DecodeFile(tomlConfFile, &conf); err != nil {
			log.Fatal("config/conf.go, toml.DecodeFile error:", err)
		}
		zapLog.Debug("load config file: ", tomlConfFile)
	})
	return conf
}

type tomlConfig struct {
	Title     string
	ApiServer *apiServer `toml:"api_server"`
	Mysql     *mysql
	Redis     *redis
	Rpc       *rpc
}
type apiServer struct {
	AppName    string
	ListenAddr string
	Debug      bool
	LogPath    string `toml:"log_path"`
}
type rpc struct {
	Ip   string
	Port string
}
type mysql struct {
	sns *dbConf
}

type dbConf struct {
	Host    string
	DbName  string `toml:"db_name"`
	User    string
	Pwd     string
	Port    string
	Charset string

	//数据库实例最大连接数312
	MaxOpenConns int `toml:"max_open_conns"`
	MaxIdleConns int `toml:"max_idle_conns"`
	//正式上线时改为3600。数据库实例上默认wait_timeout=28800, interactive_timeout=28800。
	MaxLifeTime int `toml:"max_life_time"`
}

type redis struct {
	Host string
	Pwd  string
	Port string

	MaxActiveConns    int `toml:"max_active_conns"`
	MaxIdleConns      int `toml:"max_idle_conns"`
	IdleTimeoutSecond int `toml:"idle_timeout_second"`

	PoolSize                int `toml:"pool_size"`
	DialTimeoutMillisecond  int `toml:"dial_timeout_millisecond"`
	ReadTimeoutMillisecond  int `toml:"read_timeout_millisecond"`
	WriteTimeoutMillisecond int `toml:"write_timeout_millisecond"`
}
