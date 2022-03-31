package code

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
	return b.msg
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
	FailedErr = BizErrorWithCode(Failed)
	ParamErr  = BizErrorWithCode(ParamInvalid)
	ServerErr = BizErrorWithCode(ServerError)
)
