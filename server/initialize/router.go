package initialize

import (
	"net/http"
	_ "slotsevo-admin/docs"
	"slotsevo-admin/global"
	"slotsevo-admin/middleware"
	"slotsevo-admin/router"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.New()
	Router.StaticFS(global.SlotsConfig.Local.Path, http.Dir(global.SlotsConfig.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	//global.TraceLog.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	//global.TraceLog.Info("use middleware cors")
	Router.Use(middleware.GinLogger(global.AccessLog)) // 请求日志
	Router.Use(middleware.GinRecovery(true))           // recovery panic
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.TraceLog.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitJwtRouter(PrivateGroup)    // jwt相关路由
		router.InitUserRouter(PrivateGroup)   // 注册用户路由
		router.InitMenuRouter(PrivateGroup)   // 注册menu路由
		router.InitEmailRouter(PrivateGroup)  // 邮件相关路由
		router.InitSystemRouter(PrivateGroup) // system相关路由
		router.InitCasbinRouter(PrivateGroup) // 权限相关路由
	}
	global.TraceLog.Info("router register success")
	return Router
}
