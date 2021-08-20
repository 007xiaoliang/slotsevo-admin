package rpc

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"slotsevo-admin/global"
	"slotsevo-admin/utils"
	"time"
)

var client *resty.Client

func RequestHttp(method string, url string, params map[string]interface{}, obj interface{}) (err error) {
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	url = fmt.Sprintf("%s:%s/%s", global.SlotsConfig.Rpc.Host, global.SlotsConfig.Rpc.Port, url)
	switch method {
	case "GET":
		p, err := utils.ParseMap(params)
		if err != nil {
			return err
		}
		_, _ = client.R().SetQueryParams(p).SetResult(obj).Get(url)
	case "POST":
		_, _ = client.R().SetBody(params).SetResult(obj).Post(url)
	default:
		err = errors.New("unsupported request method")
	}
	return
}

func init() {
	client = resty.New()
	// Set client timeout as per your need
	client.SetTimeout(time.Duration(global.SlotsConfig.Rpc.Timeout) * time.Minute)

	client.SetRetryCount(global.SlotsConfig.Rpc.RetryCount)
	client.SetRetryWaitTime(time.Duration(global.SlotsConfig.Rpc.RetryWaitTime) * time.Second)
	client.SetRetryMaxWaitTime(time.Duration(global.SlotsConfig.Rpc.RetryMaxWaitTime) * time.Second)
	client.SetContentLength(true)

	//middleware
	//client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
	//	global.BaseSendLog.Info("send request", zap.String("url", req.URL))
	//	return nil
	//})
	//client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
	//	global.BaseSendLog.Info("receive data", zap.Any("status", resp.Status()))
	//	return nil
	//})
}
