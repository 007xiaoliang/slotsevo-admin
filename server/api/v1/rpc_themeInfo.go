package v1

import (
	"github.com/gin-gonic/gin"
	"slotsevo-admin/model/request"
	"slotsevo-admin/model/response"
	"slotsevo-admin/rpc"
)

func GetThemeInfo(c *gin.Context) {
	var themeInfo request.ThemeInfo
	if err := c.ShouldBindJSON(&themeInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var rpcThemeInfo [][]string
	err := rpc.RequestHttpWithDetail("POST", "debug", themeInfo.RpcHost, themeInfo.RpcPort, map[string]interface{}{
		"cmd": "get_debug_themes",
	}, &rpcThemeInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(rpcThemeInfo, "获取成功", c)
}
