package router

import (
	"github.com/gin-gonic/gin"
	"slotsevo-admin/api/v1"
)

func InitRpcRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("rpc").Use()
	{
		BaseRouter.POST("userInfo", v1.GetUserInfo)
		BaseRouter.POST("themeInfo", v1.GetThemeInfo)
		BaseRouter.POST("actInfo", v1.GetActInfo)
	}
	return BaseRouter
}
