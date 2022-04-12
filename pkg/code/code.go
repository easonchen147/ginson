package code

const Success = 0 // 正常

// 系统错误
const (
	ServerFailed = 10000 // 服务器异常
	MysqlFailed  = 10001 // mysql访问异常
	RedisFailed  = 10002 // redis访问异常
	MongoFailed  = 10003 // mongo访问异常
	KafkaFailed  = 10004 // kafka访问异常
)

// 业务错误
const (
	Failed      = 50000
	LoginFailed = 50001
)

// 参数校验相关
const (
	ParamInvalid  = 40000
	TokenInvalid  = 40001
	TokenEmpty    = 40002
	OpenIdInvalid = 40003
)

var codeToMsg = map[int]string{
	ServerFailed:  "服务器内部错误，请稍后再试",
	Failed:        "业务异常",
	ParamInvalid:  "参数错误",
	LoginFailed:   "登录失败",
	TokenInvalid:  "Token不合法",
	TokenEmpty:    "Token为空",
	OpenIdInvalid: "OpenId为空",
	MysqlFailed:   "数据库异常",
	RedisFailed:   "缓存库异常",
}
