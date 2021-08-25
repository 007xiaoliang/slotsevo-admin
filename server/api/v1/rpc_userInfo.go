package v1

import (
	"github.com/gin-gonic/gin"
	"slotsevo-admin/model/request"
	"slotsevo-admin/model/response"
	"slotsevo-admin/rpc"
	"slotsevo-admin/rpc/model"
	"slotsevo-admin/utils"
)

// @Tags Base
// @Summary rpc查询用户信息
// @Produce  application/json
// @Param data body request.UserInfo true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/login [post]
func GetUserInfo(c *gin.Context) {
	var userInfo request.UserInfo
	_ = c.ShouldBindJSON(&userInfo)
	if err := utils.Verify(userInfo, utils.UserIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var rpcUserInfo  model.RpcUserInfo
	err := rpc.RequestHttp("POST", "debug", map[string]interface{}{
		"cmd": "get_user_baseinfo",
		"params":map[string]string{
			"userid": userInfo.UserID,
		},
	}, &rpcUserInfo)
	if err!=nil{
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(rpcUserInfo, "获取成功", c)
}
