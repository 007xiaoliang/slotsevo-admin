package router

import (
	"github.com/gin-gonic/gin"
	"slotsevo-admin/api/v1"
	"slotsevo-admin/middleware"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("createApi", v1.CreateApi)               // 创建Api
		ApiRouter.POST("deleteApi", v1.DeleteApi)               // 删除Api
		ApiRouter.POST("getApiList", v1.GetApiList)             // 获取Api列表
		ApiRouter.POST("getApiById", v1.GetApiById)             // 获取单条Api消息
		ApiRouter.POST("updateApi", v1.UpdateApi)               // 更新api
		ApiRouter.POST("getAllApis", v1.GetAllApis)             // 获取所有api
		ApiRouter.DELETE("deleteApisByIds", v1.DeleteApisByIds) // 删除选中api
	}
}
