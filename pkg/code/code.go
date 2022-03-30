package code

// 系统使用
const (
	Success      = 0  //正常
	Failed       = -1 //失败
	ParamInvalid = -2 //参数错误
)

// 业务使用
const (
	LoginFailed   = 10000
	TokenInvalid  = 10001
	OpenIdInvalid = 10002
	TokenEmpty    = 10003
)

var codeToMsg = map[int]string{
	Failed:        "服务器异常",
	ParamInvalid:  "参数错误",
	LoginFailed:   "登录失败",
	TokenInvalid:  "Token不合法",
	OpenIdInvalid: "OpenId为空",
}
