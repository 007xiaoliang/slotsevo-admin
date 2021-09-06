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
		}
	}()
	var actInfo request.ActInfo
	if err := c.ShouldBindJSON(&actInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var wg sync.WaitGroup
	var responseChan = make(chan interface{}, 20)
	wg.Add(2)
	go func() {
		defer wg.Done()
		var rpcActivityType *[]string
		err := rpc.RequestHttpWithDetail("POST", "debug", actInfo.RpcHost, actInfo.RpcPort, map[string]interface{}{
			"cmd": "get_activity_type",
		}, rpcActivityType)
		if err != nil {
			responseChan <- err
		} else {
			responseChan <- rpcActivityType
		}
	}()
	go func() {
		defer wg.Done()
		var rpcServerTime string
		err := rpc.RequestHttpWithDetail("POST", "debug", actInfo.RpcHost, actInfo.RpcPort, map[string]interface{}{
			"cmd": "get_server_time",
		}, &rpcServerTime)
		if err != nil {
			responseChan <- err
		} else {
			responseChan <- rpcServerTime
		}
	}()

	// 等待所以协程执行完毕
	wg.Wait() // 当计数器为0时, 不再阻塞
	// 关闭接收结果channel
	close(responseChan)
	var rpcServerTime string
	var rpcActivityType *[][]string
	for value := range responseChan {
		switch t := value.(type) {
		case error:
			response.FailWithMessage(t.Error(), c)
			return
		case string:
			rpcServerTime = value.(string)
		case *[][]string:
			rpcActivityType = value.(*[][]string)
		}
	}
	activityTypeResponse := &model.ActivityTypeResponse{
		ServerTime:       rpcServerTime,
		ActivityTypeList: rpcActivityType,
	}
	response.OkWithDetailed(activityTypeResponse, "获取成功", c)
}
