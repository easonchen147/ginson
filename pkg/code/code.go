package code

// 系统使用
const (
	Success      = 0  //正常
	Failed       = -1 //失败
	ParamInvalid = -2 //参数错误
)

// 业务使用
const (
	EmailExists     = 10000
	EmailNotExists  = 10001
	EmailOrPswWrong = 10002
	LoginFailed     = 10003
)

var codeToMsg = map[int]string{
	Failed:          "服务器异常",
	ParamInvalid:    "参数错误",
	EmailExists:     "邮箱已存在",
	EmailNotExists:  "邮箱不存在",
	EmailOrPswWrong: "邮箱或密码错误",
	LoginFailed:     "登录失败",
}
