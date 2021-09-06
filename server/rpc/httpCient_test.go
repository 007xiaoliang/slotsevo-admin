package rpc

import (
	"fmt"
	"testing"
)

func TestRequestHttp(t *testing.T) {
	var activityTypeArr []string
	err := RequestHttp("POST", "debug", map[string]interface{}{
		"cmd": "get_debug_themes",
	}, &activityTypeArr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(activityTypeArr)
}

func TestRequestHttpWithDetail(t *testing.T) {
	//var rpcActivityType string
	resp, err := RequestHttpWithDetail("POST", "debug", "192.168.2.34", "7001", map[string]interface{}{
		"cmd": "get_server_time",
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.String())
}
