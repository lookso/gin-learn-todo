package db

import (
	"database/sql"
	"gin-learn-todo/setting"
	"github.com/jinzhu/gorm"
)

func New() (map[string]*RWManager, error) {
	clusters, err := Build()
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

func Read(name ...string) *gorm.DB {
	var key string
	if len(name) == 0 {
		key = setting.Conf.Db.Default
	} else {
		key = name[0]
	}

	if c, ok := clusters[key]; ok {
		return c.Read()
	} else {
		panic("package pkg/db is not init")
	}
	return nil
}

func Write(name ...string) *gorm.DB {
	var key string
	if len(name) == 0 {
		key = setting.Conf.Db.Default
	} else {
		key = name[0]
	}

	if c, ok := clusters[key]; ok {
		return c.Write()
	} else {
		panic("package pkg/db is not init")
	}
	return nil
}

func Stats() map[string]map[string]sql.DBStats {
	stats := make(map[string]map[string]sql.DBStats)
	for k, v := range clusters {
		stats[k] = v.Stats()
	}
	return stats
}

func Close() {
	if clusters != nil {
		for _, v := range clusters {
			v.Close()
		}
	}
}
