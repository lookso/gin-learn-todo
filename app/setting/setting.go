package setting

import (
	"flag"
	"gin-learn-todo/app/public/zaplog"
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
	"sync"
)

// 所有配置
var Conf *tomlConf

var tomlConfLoadOnce sync.Once

//调用前，需先在外部初始化
var tomlConfFile string

func init() {
	var confPointer = flag.String("conf", "../gin-learn-todo/app/config/dev.toml", "-conf ../gin-learn-todo/app/config/dev.toml")
	flag.Parse()
	configFile, err := filepath.Abs(*confPointer)
	if err != nil {
		log.Fatal("get config file err: ", err)
	}
	tomlConfFile = configFile
}

func Init() *tomlConf {
	// flag 放在该方法里会导致如果多次 调用Config()方法的时候导致flag redefined flag被重复调用到,导致重读定义了
	// 所以单独拿出去

	// 使用了sync.Once的Do方法, 当且仅当第一次被调用时才执行函数
	tomlConfLoadOnce.Do(func() {
		if _, err := toml.DecodeFile(tomlConfFile, &Conf); err != nil {
			log.Fatal("config/conf.go, toml.DecodeFile error:", err)
		}
		log.Fatal("load config file: ", tomlConfFile)
	})
	return Conf
}

type tomlConf struct {
	Title string

	ApiServer *apiServer     `toml:"api_server"`
	LogConfig *zaplog.Config `toml:"log_config"`
	Mysql     *Mysql        `toml:"mysql"`
	Redis     *Redis         `toml:"redis"`
}

type apiServer struct {
	AppName    string `toml:"app_name"`
	ListenAddr string `toml:"listen_addr"`
	Debug      int    `toml:"debug"`

	// 时区格式：UTC、PRC、Asia/Calcutta 、Asia/Kolkata
	Timezone string `toml:"timezone"`
}

type Mysql struct {
	Sns *DbConf
}

type DbConf struct {
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

type Redis struct {
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
