package initial

import (
	"context"
	"entgo.io/ent/dialect/sql"
	appClient "github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen"
	appMigrate "github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/migrate"
	"github.com/wangxg422/XishangOS-backend/common/constant"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/middleware/logger"
	"go.uber.org/zap"
)

var AppDbClient *appClient.Client

func CreateMySQLClient(drv *sql.Driver) {
	client := appClient.NewClient(appClient.Driver(drv))

	if global.AppConfig.App.Mode == constant.DEBUG {
		AppDbClient = client.Debug()
	} else {
		AppDbClient = client
	}

	if err := AppDbClient.Schema.Create(context.Background(), appMigrate.WithForeignKeys(false)); err != nil {
		logger.Error("failed creating schema resources: %v", zap.String("error", err.Error()))
		panic(err)
	}
}
