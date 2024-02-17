package main

import (
	"github.com/wangxg422/XishangOS-backend/initial"
)

func main() {
	initial.Viper()
	initial.InitLog()
	initial.InitDb()
	initial.Redis()
	initial.InitRouter()
}
