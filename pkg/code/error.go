package code

import "fmt"

type BizCodeMsg interface {
	Code() int
	Msg() string
}

type BizErr interface {
	BizCodeMsg
	error
}

type bizError struct {
	code int
	msg  string
}

func (b *bizError) Code() int {
	return b.code
}

func (b *bizError) Msg() string {
	if b.msg != "" {
		return b.msg
	}
	return codeToMsg[b.code]
}

func (b *bizError) Error() string {
	return fmt.Sprintf("%d %s", b.Code(), b.Msg())
}

func BizErrorWithCode(code int) BizErr {
	return &bizError{code: code}
}

func BizErrorWithCodeMsg(code int, msg string) BizErr {
	return &bizError{code: code, msg: msg}
}

func BizError(err error) BizErr {
	return &bizError{code: Failed, msg: err.Error()}
}

// 系统常用error
var (
	FailedError = BizErrorWithCode(Failed)
	ParamError  = BizErrorWithCode(ParamInvalid)
	ServerError = BizErrorWithCode(ServerFailed)
	MysqlError  = BizErrorWithCode(MysqlFailed)
	RedisError  = BizErrorWithCode(RedisFailed)
)

// 定义模块功能错误
var (
	LoginFailedError   = BizErrorWithCode(LoginFailed)
	TokenInvalidError  = BizErrorWithCode(TokenInvalid)
	TokenEmptyError    = BizErrorWithCode(TokenEmpty)
	OpenIdInvalidError = BizErrorWithCode(OpenIdInvalid)
)
