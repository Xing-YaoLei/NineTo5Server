package main

import (
	"NineTo5Server/core"
	"NineTo5Server/global"
	"NineTo5Server/initialize"
	"go.uber.org/zap"
	"os"
)

func main() {
	global.NineTo5_VP = core.Viper()
	global.NineTo5_LOG = core.Zap() // 初始化zap日志
	zap.ReplaceGlobals(global.NineTo5_LOG)
	global.NineTo5_DB = initialize.Gorm() // gorm连接数据库
	if global.NineTo5_DB == nil {
		global.NineTo5_LOG.Error("mysql连接失败，退出程序")
		os.Exit(127)
	} else {
		initialize.RegisterTables(global.NineTo5_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.NineTo5_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
