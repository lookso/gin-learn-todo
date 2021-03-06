package setting

import (
	"flag"
	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"log"
	"path/filepath"
	"sync"
)

type tomlConf struct {
	Title  string
	App    *app    `toml:"app"`
	Db     *db     `toml:"db"`
	Redis  *redis  `toml:"redis"`
	Sentry *sentry `toml:"sentry"`
	Logger *logger `toml:"logger"`
}

type app struct {
	Addr  string `toml:"addr"`
	Debug bool   `toml:"debug"`
	// 时区格式：UTC、PRC、Asia/Calcutta 、Asia/Kolkata
	Timezone string `toml:"timezone"`
	AppId    string `toml:"app_id"`
	Env      string `toml:"env"`
}

type db struct {
	Default  string                    `toml:"default"` // 默认集群
	Clusters map[string]ClustersConfig `json:"clusters" toml:"clusters"`
}
type ClustersConfig struct {
	Master Endpoint   `toml:"master"`
	Slaves []Endpoint `toml:"slaves"`
}
type Endpoint struct {
	Dns          string `toml:"dns"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
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
	Env string `toml:env`
	Dsn string `toml:"dsn"`
}

// Config 日志的可配置项
type logger struct {
	Level zap.AtomicLevel `json:"level"`
	// Encoding json or console
	Encoding string `json:"encoding,omitempty"`
	// Path 输出目录，支持文件，stdout stderr 等
	//Path []string `json:"outputs,omitempty"`
	Path string `json:"path"`
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
