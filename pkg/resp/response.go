package resp

import "ginson/pkg/code"

// CommonResponse api响应统一定义
type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponseSuccess(data interface{}) *CommonResponse {
	return &CommonResponse{
		Code: code.Success,
		Msg:  "ok",
		Data: data,
	}
}

func NewResponseFailed() *CommonResponse {
	return &CommonResponse{
		Code: code.FailedErr.Code(),
		Msg:  code.FailedErr.Msg(),
		Data: nil,
	}
}

func NewResponseFailedBinding(bindErr error) *CommonResponse {
	return &CommonResponse{
		Code: code.ParamInvalid,
		Msg:  bindErr.Error(),
		Data: nil,
	}
}

func NewResponseFailedMsg(msg string) *CommonResponse {
	return &CommonResponse{
		Code: code.Failed,
		Msg:  msg,
		Data: nil,
	}
}

func NewResponseFailedCodeMsg(code int, msg string) *CommonResponse {
	return &CommonResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func NewResponseFailedBizErr(bizErr code.BizErr) *CommonResponse {
	return &CommonResponse{
		Code: bizErr.Code(),
		Msg:  bizErr.Msg(),
		Data: nil,
	}
}

func NewResponseFailedErr(err error) *CommonResponse {
	return &CommonResponse{
		Code: code.Failed,
		Msg:  err.Error(),
		Data: nil,
	}
}
