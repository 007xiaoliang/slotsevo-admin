package rpc

import (
	"fmt"
	"slotsevo-admin/core"
	"slotsevo-admin/global"
	"testing"
)

func TestRequestHttp(t *testing.T) {
	global.SlotsVp = core.Viper("../config.yaml")
	global.TraceLog, global.AccessLog, global.BaseSendLog, global.ErrorLog = core.Zap()
	var activityTypeArr []string
	err := RequestHttp("POST", "debug", map[string]interface{}{
		"cmd": "get_debug_themes",
	}, &activityTypeArr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(activityTypeArr)
}
