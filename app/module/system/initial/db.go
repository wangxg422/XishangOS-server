package initial

import (
	"context"
	"entgo.io/ent/dialect/sql"
	sysClient "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	sysMigrate "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/migrate"
	"github.com/wangxg422/XishangOS-backend/common/constant"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/middleware/logger"
	"go.uber.org/zap"
)

var SysDbClient *sysClient.Client

func CreateMySQLClient(drv *sql.Driver) {
	client := sysClient.NewClient(sysClient.Driver(drv))

	if global.AppConfig.App.Mode == constant.DEBUG {
		SysDbClient = client.Debug()
	} else {
		SysDbClient = client
	}

	if err := SysDbClient.Schema.Create(context.Background(), sysMigrate.WithForeignKeys(false)); err != nil {
		logger.Error("failed creating schema resources: %v", zap.String("error", err.Error()))
		panic(err)
	}
}
