/*
@Time : 2019-11-17 17:05 
@Author : Tenlu
@File : conf
@Software: GoLand
*/
package conf

import (
	"flag"
	"fmt"
	zapLog "gin-learn-todo/app/utils/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var conf *Config
var yamlConfFile string

type Config struct {
	Title  string
	Debug  bool
	Listen string
	Log    zapLog.Config
	Rpc    Rpc
}
type Rpc struct {
	Port string
}

func (_ Config) InitConfig() *Config {
	var confPointer = flag.String("conf", "../gin-learn-todo/app/config/app.yaml", "-conf ../gin-learn-todo/app/config/app.yaml")
	flag.Parse()

	configFile, err := filepath.Abs(*confPointer)
	if err != nil {
		log.Println("get-conf-file", err)
	}
	fmt.Println(configFile)
	yamlConfFile = configFile

	var data []byte

	// 加载config，目前支持了yaml格式
	if data, err = ioutil.ReadFile(yamlConfFile); err != nil {
		fmt.Println("config file read err:", err)
		os.Exit(1)
	}

	if err = yaml.Unmarshal(data, &conf); err != nil {
		fmt.Println("config file unmarshal err:", err)
		os.Exit(1)
	}
	return conf
}
