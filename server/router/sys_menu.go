package router

import (
	"github.com/gin-gonic/gin"
	"slotsevo-admin/api/v1"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu")
	{
		MenuRouter.POST("getMenu", v1.GetMenu)         // 获取菜单树
		MenuRouter.POST("getMenuList", v1.GetMenuList) // 分页获取基础menu列表
	}
	return MenuRouter
}
