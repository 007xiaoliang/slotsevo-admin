package router

import (
	"github.com/gin-gonic/gin"
	"slotsevo-admin/api/v1"
	"slotsevo-admin/middleware"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
