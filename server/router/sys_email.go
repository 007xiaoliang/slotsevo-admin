package router

import (
	"github.com/gin-gonic/gin"
	"slotsevo-admin/api/v1"
	"slotsevo-admin/middleware"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		UserRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
