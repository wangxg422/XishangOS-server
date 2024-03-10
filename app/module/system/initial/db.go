package initial

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/migrate"
	_ "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/runtime"
	"github.com/wangxg422/XishangOS-backend/common/constant"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/initial/logger"
	"go.uber.org/zap"
)

var SysDbClient *codegen.Client

func CreateMySQLClient(drv *sql.Driver) {
	client := codegen.NewClient(codegen.Driver(drv))

	if global.AppConfig.App.Mode == constant.DEBUG {
		SysDbClient = client.Debug()
	} else {
		SysDbClient = client
	}

	if err := SysDbClient.Schema.Create(context.Background(), migrate.WithForeignKeys(global.AppConfig.Database.EnableForeignKey)); err != nil {
		logger.Error("failed creating schema resources: %v", zap.String("error", err.Error()))
		panic(err)
	}
}
