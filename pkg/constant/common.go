package constant

import "ginson/pkg/conf"

const (
	TraceIdKey = "TraceId"
)

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

const (
	OauthSourceWxMini = "WxMini"
	OauthSourceWechat = "Wechat"
	OauthSourceQQ     = "QQ"
	OauthSourceQQMini = "QQMini"
)

var (
	TokenSecret = []byte(conf.AppConf.Ext.TokenSecret)
)

