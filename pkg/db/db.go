package db

import (
	"errors"
	"gin-learn-todo/pkg/log"
	"gin-learn-todo/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	ErrClustersEmpty = errors.New("clusters is empty")
)

var (
	clusters map[string]*RWManager
)

// Init 根据Config初始化
func Init() error {
	var err error
	clusters, err = Build()
	if err != nil {
		return err
	}
	return nil
}

func Build() (map[string]*RWManager, error) {
	c := setting.Conf.Db
	if len(c.Clusters) == 0 {
		log.Sugar().Errorf("clusters len empty %v", ErrClustersEmpty)
		return nil, ErrClustersEmpty
	}
	clusters = make(map[string]*RWManager, len(c.Clusters))
	for name, clusterConfig := range c.Clusters {
		var err error
		clusters[name], err = buildCluster(clusterConfig)
		if err != nil {
			log.Sugar().Errorf("Build err(%v)", err)
			return nil, err
		}
	}
	return clusters, nil
}

func buildCluster(c setting.ClustersConfig) (*RWManager, error) {
	master, err := buildEndpoint(c.Master)
	if err != nil {
		log.Sugar().Errorf("buildCluster err(%v)", err)
		return nil, err
	}
	var slaves []*gorm.DB
	for _, v := range c.Slaves {
		if db, err := buildEndpoint(v); err != nil {
			log.Sugar().Errorf("buildCluster slaves err(%v)", err)
			return nil, err
		} else {
			slaves = append(slaves, db)
		}
	}
	return &RWManager{
		master: master,
		slaves: slaves,
	}, nil
}

func buildEndpoint(c setting.Endpoint) (*gorm.DB, error) {
	if db, err := gorm.Open("mysql", c.Dns); err != nil {
		log.Sugar().Errorf("buildEndpoint gorm.Open err(%v)", err)
		return nil, err
	} else {
		if c.MaxOpenConns > 0 {
			db.DB().SetMaxOpenConns(c.MaxOpenConns)
		}
		if c.MaxIdleConns > 0 {
			db.DB().SetMaxIdleConns(c.MaxIdleConns)
		}
		if c.MaxLifeTime > 0 {
			db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifeTime) * time.Second)
		}
		if setting.Conf.App.Env != "prod" {
			db.LogMode(true)
		}
		return db, nil
	}
}
