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
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.SlotsLog.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.SlotsLog.Error(s.ListenAndServe().Error())
}
