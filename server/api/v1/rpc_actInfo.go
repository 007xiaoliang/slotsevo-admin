package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slotsevo-admin/global"
	"slotsevo-admin/model/request"
	"slotsevo-admin/model/response"
	"slotsevo-admin/rpc"
	"slotsevo-admin/rpc/model"
	"sync"
)

func GetActInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			global.ErrorLog.Info("[Recovery from panic]", zap.Any("err", err))
			response.FailWithMessage(err.(string), c)
		}
	}()
	var actInfo request.ActInfo
	if err := c.ShouldBindJSON(&actInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var wg sync.WaitGroup
	var responseChan = make(chan interface{}, 2)
	var rpcActivityType []string
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := rpc.RequestHttpWithDetail("POST", "debug", actInfo.RpcHost, actInfo.RpcPort, map[string]interface{}{
			"cmd": "get_activity_type",
		}, &rpcActivityType)
		if err != nil {
			responseChan <- err
		}
	}()
	go func() {
		defer wg.Done()
		resp, err := rpc.RequestHttpWithDetail("POST", "debug", actInfo.RpcHost, actInfo.RpcPort, map[string]interface{}{
			"cmd": "get_server_time",
		}, nil)
		if err != nil {
			responseChan <- err
		} else {
			rpcServerTime := resp.String()
			responseChan <- rpcServerTime
		}
	}()
	wg.Wait()
	close(responseChan)
	var rpcServerTime string
	for value := range responseChan {
		switch t := value.(type) {
		case error:
			response.FailWithMessage(t.Error(), c)
			return
		case string:
			rpcServerTime = value.(string)
		}
	}
	activityTypeResponse := &model.ActivityTypeResponse{
		ServerTime:       rpcServerTime,
		ActivityTypeList: rpcActivityType,
	}
	response.OkWithDetailed(activityTypeResponse, "获取成功", c)
}
