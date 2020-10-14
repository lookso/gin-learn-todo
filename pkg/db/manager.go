package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
)

type RWManager struct {
	master *gorm.DB
	slaves []*gorm.DB
}

func (m *RWManager) Write() *gorm.DB {
	return m.master
}

func (m *RWManager) Read() *gorm.DB {
	l := len(m.slaves)
	if l > 0 {
		// 使用了全局的rand，理论上性能会有些差，但是考虑到db的速度就差别不大了，后面换第三方的吧
		k := rand.Intn(l)
		return m.slaves[k]
	}
	return m.master
}

func (m *RWManager) Stats() map[string]sql.DBStats {
	stats := make(map[string]sql.DBStats)
	stats["master"] = m.master.DB().Stats()
	for k, v := range m.slaves {
		stats["slave"+strconv.Itoa(k)] = v.DB().Stats()
	}
	return stats
}

func (m *RWManager) Close() {
	if err := m.master.Close(); err != nil {
		fmt.Println("master close err", err)
	}
	for _, v := range m.slaves {
		err := v.Close()
		fmt.Println("slave close err", err)
	}
}
