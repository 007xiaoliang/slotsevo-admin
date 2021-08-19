package main

import (
	"slotsevo-admin/core"
	"slotsevo-admin/global"
	"slotsevo-admin/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.SlotsVp = core.Viper()                                                       // 初始化Viper
	global.TraceLog, global.AccessLog, global.BaseSendLog, global.ErrorLog = core.Zap() // 初始化zap日志库
	global.SlotsDb = initialize.Gorm()                                                  // gorm连接并初始化数据库
	initialize.Timer()                                                                  // 初始化任务
	if global.SlotsDb != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.SlotsDb.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
