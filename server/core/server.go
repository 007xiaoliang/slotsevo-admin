package core

import (
	"fmt"
	"go.uber.org/zap"
	"slotsevo-admin/global"
	"slotsevo-admin/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.SlotsConfig.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.SlotsConfig.System.Addr)
	s := initServer(address, Router)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.TraceLog.Info("server run success on ", zap.String("address", address))
	global.TraceLog.Error(s.ListenAndServe().Error())
}
