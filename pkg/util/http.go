package util

import (
	"context"
	"ginson/conf"
	"github.com/go-resty/resty/v2"
	"time"
)

var httpClient *resty.Client

func init() {
	httpClient = resty.New()
	httpClient.SetTimeout(time.Second * 5)
	httpClient.SetDebug(conf.AppConf.IsDevEnv())
}

// GetHttpClient 获取http client 实例
func GetHttpClient(ctx context.Context) *resty.Client {
	return httpClient
}

// Post 快速发起post请求
func Post(ctx context.Context, url string, data interface{}, result interface{}) error {
	_, err := httpClient.R().SetBody(data).SetResult(&result).Post(url)
	return err
}

// Get  快速发起get请求
func Get(ctx context.Context, url string, result interface{}) error {
	_, err := httpClient.R().SetResult(&result).Get(url)
	return err
}
