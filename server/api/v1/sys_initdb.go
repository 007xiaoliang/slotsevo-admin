package v1

import (
	"go.uber.org/zap"
	"slotsevo-admin/global"
	"slotsevo-admin/model/request"
	"slotsevo-admin/model/response"
	"slotsevo-admin/service"

	"github.com/gin-gonic/gin"
)

// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /init/initdb [post]
func InitDB(c *gin.Context) {
	if global.GvaDb != nil {
		global.GvaLog.Error("非法访问")
		response.FailWithMessage("非法访问", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GvaLog.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := service.InitDB(dbInfo); err != nil {
		global.GvaLog.Error("自动创建数据库失败", zap.Any("err", err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志", c)
		return
	}
	response.OkWithData("自动创建数据库成功", c)
}

// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /init/checkdb [post]
func CheckDB(c *gin.Context) {
	if global.GvaDb != nil {
		global.GvaLog.Info("数据库无需初始化")
		response.OkWithDetailed(gin.H{
			"needInit": false,
		}, "数据库无需初始化", c)
		return
	} else {
		global.GvaLog.Info("前往初始化数据库")
		response.OkWithDetailed(gin.H{
			"needInit": true,
		}, "前往初始化数据库", c)
		return
	}
}
