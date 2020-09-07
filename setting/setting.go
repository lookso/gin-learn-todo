package setting

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
	"sync"
)

type tomlConf struct {
	Title string

	ApiServer *apiServer `toml:"api_server"`
	Mysql     *mysql     `toml:"mysql"`
	Redis     *redis     `toml:"redis"`
	Sentry    *sentry    `toml:"sentry"`
}

type apiServer struct {
	ListenAddr string `toml:"listen_addr"`
	Debug      bool   `toml:"debug"`
	// 时区格式：UTC、PRC、Asia/Calcutta 、Asia/Kolkata
	Timezone string `toml:"timezone"`
	AppId    string `toml:"app_id"`
}

type mysql struct {
	Sns *dbConf
}

type dbConf struct {
	Addr         string `toml:"addr"`           // 连接信息
	MaxOpenConns int    `toml:"max_open_conns"` // 用于设置最大打开的连接数，默认值为0表示不限制
	MaxIdleConns int    `toml:"max_idle_conns"` // 用于设置闲置的连接数
	MaxLifeTime  int    `toml:"max_life_time"`
}

type redis struct {
	Addr     string `toml:"addr"`
	Passport string `toml:"passport"`
	Db       int    `toml:"db"`

	MaxActiveConns          int `toml:"max_active_conns"`
	MaxIdleConns            int `toml:"max_idle_conns"`
	IdleTimeoutSecond       int `toml:"idle_timeout_second"`
	PoolSize                int `toml:"pool_size"`
	DialTimeoutMillisecond  int `toml:"dial_timeout_millisecond"`
	ReadTimeoutMillisecond  int `toml:"read_timeout_millisecond"`
	WriteTimeoutMillisecond int `toml:"write_timeout_millisecond"`
}

type sentry struct {
	Dsn string `toml:"dsn"`
}

// 所有配置
var Conf *tomlConf

var tomlConfLoadOnce sync.Once

func init() {
	var confPointer = flag.String("conf", "./config/dev.toml", "-conf ./config/dev.toml")
	flag.Parse()
	configFile, err := filepath.Abs(*confPointer)
	if err != nil {
		log.Fatal("get config file err: ", err)
	}

	// 使用了sync.Once的Do方法, 当且仅当第一次被调用时才执行函数
	tomlConfLoadOnce.Do(func() {
		if _, err := toml.DecodeFile(configFile, &Conf); err != nil {
			log.Fatal("toml.DecodeFile error:", err)
		}
		log.Println("load config file success:", configFile)
	})
}
