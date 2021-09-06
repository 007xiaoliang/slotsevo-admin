package rpc

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"slotsevo-admin/global"
	"slotsevo-admin/utils"
	"time"
)

var client *resty.Client

func RequestHttp(method string, url string, params map[string]interface{}, obj interface{}) (err error) {
	url = fmt.Sprintf("%s:%s/%s", global.SlotsConfig.Rpc.Host, global.SlotsConfig.Rpc.Port, url)
	switch method {
	case "GET":
		p, err := utils.ParseMap(params)
		if err != nil {
			return err
		}
		_, err = client.R().SetQueryParams(p).SetResult(obj).Get(url)
	case "POST":
		_, err = client.R().SetBody(params).SetResult(obj).Post(url)
	default:
		err = errors.New("unsupported request method")
	}
	return
}

func RequestHttpWithDetail(method, url, host, port string, params map[string]interface{}, obj interface{}) (resp *resty.Response, err error) {
	url = fmt.Sprintf("http://%s:%s/%s", host, port, url)
	switch method {
	case "GET":
		p, err := utils.ParseMap(params)
		if err != nil {
			return nil, err
		}
		_, err = client.R().SetQueryParams(p).SetResult(obj).Get(url)
	case "POST":
		if obj != nil {
			resp, err = client.R().SetBody(params).SetResult(obj).Post(url)
		} else {
			resp, err = client.R().SetBody(params).Post(url)
		}
	default:
		err = errors.New("unsupported request method")
	}
	return resp, err
}

func init() {
	client = resty.New()
	// Set client timeout as per your need
	client.SetTimeout(time.Duration(global.SlotsConfig.Rpc.Timeout) * time.Minute)
	client.SetRetryCount(global.SlotsConfig.Rpc.RetryCount)
	client.SetRetryWaitTime(time.Duration(global.SlotsConfig.Rpc.RetryWaitTime) * time.Second)
	client.SetRetryMaxWaitTime(time.Duration(global.SlotsConfig.Rpc.RetryMaxWaitTime) * time.Second)
	client.SetContentLength(true)
	// 并发情况下设置header可能会引起concurrent write map错误。因此初始化时将map的设置都初始化
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	//middleware
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		global.BaseSendLog.Info("send request", zap.String("url", req.URL))
		return nil
	})
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		global.BaseSendLog.Info("receive data", zap.Any("status", resp.Status()))
		return nil
	})
}
