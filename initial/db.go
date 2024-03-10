package initial

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	appInitial "github.com/wangxg422/XishangOS-backend/app/module/application/initial"
	sysInitial "github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	sysDbClient "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	sysMigrate "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/migrate"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/initial/logger"
	"go.uber.org/zap"
	"time"
)

func InitDb() {
	if global.AppConfig.Database.Type == "mysql" {
		InitMySQLWithSqlDb()
	}
}

func InitMySQL() {
	dbConfig := global.AppConfig.Database.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.Address, dbConfig.Port, dbConfig.Dbname, dbConfig.ConnConfig)

	client, err := sysDbClient.Open("mysql", dsn)
	if err != nil {
		logger.Error("failed opening connection to mysql: %v", zap.String("err", err.Error()))
	}

	if err := client.Debug().Schema.Create(context.Background(), sysMigrate.WithForeignKeys(false)); err != nil {
		logger.Error("failed creating schema resources: %v", zap.String("err", err.Error()))
	}
}

func InitMySQLWithSqlDb() {
	dbConfig := global.AppConfig.Database.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.Address, dbConfig.Port, dbConfig.Dbname, dbConfig.ConnConfig)

	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	db.SetMaxIdleConns(dbConfig.MaxIdleConn)
	db.SetMaxOpenConns(dbConfig.MaxOpenConn)
	db.SetConnMaxLifetime(time.Hour)

	// 初始化各应用模块DB客户端
	sysInitial.CreateMySQLClient(drv)
	appInitial.CreateMySQLClient(drv)
}
