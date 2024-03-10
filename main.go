package main

import (
	"github.com/wangxg422/XishangOS-backend/initial"
	"github.com/wangxg422/XishangOS-backend/initial/logger"
)

func main() {
	initial.Viper()
	logger.InitLog()
	initial.InitDb()
	initial.Redis()
	initial.InitRouter()
}
