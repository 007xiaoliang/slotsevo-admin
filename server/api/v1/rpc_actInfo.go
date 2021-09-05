package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"slotsevo-admin/model/request"
	"slotsevo-admin/model/response"
	"slotsevo-admin/rpc"
	"slotsevo-admin/rpc/model"
	"sync"
)

func GetActInfo(c *gin.Context) {
	var actInfo request.ActInfo
	if err := c.ShouldBindJSON(&actInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var rpcActivityType []string
	var rpcServerTime string
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := rpc.RequestHttpWithDetail("POST", "debug", actInfo.RpcHost, actInfo.RpcPort, map[string]interface{}{
			"cmd": "get_activity_type",
		}, &rpcActivityType)
		fmt.Print("111111111")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}()
	go func() {
		defer wg.Done()
		err := rpc.RequestHttpWithDetail("POST", "debug", actInfo.RpcHost, actInfo.RpcPort, map[string]interface{}{
			"cmd": "get_server_time",
		}, &rpcServerTime)
		fmt.Print("222222222")
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}()
	wg.Wait()
	activityTypeResponse := &model.ActivityTypeResponse{
		ServerTime: rpcServerTime,
		ActivityTypeList: rpcActivityType,
	}
	response.OkWithDetailed(activityTypeResponse, "获取成功", c)
}
