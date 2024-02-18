package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	entadapter "github.com/casbin/ent-adapter"
	"github.com/wangxg422/XishangOS-backend/global"
	"sync"
)

type AppCasbin struct {
}

var AppCasbinService = new(AppCasbin)

// 持久化到数据库
var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (m *AppCasbin) GetCasbinEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		dbConfig := global.AppConfig.Database.Mysql
		dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.Address, dbConfig.Port, dbConfig.Dbname, dbConfig.ConnConfig)

		a, _ := entadapter.NewAdapter("mysql", dsn)
		// 使用ent client作为适配器
		syncedEnforcer, _ = casbin.NewSyncedEnforcer("resource/casbin/rbac_model.conf", a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
