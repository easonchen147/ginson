package resp

import "ginson/pkg/code"

type CommonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponseSuccess(data interface{}) *CommonResp {
	return &CommonResp{
		Code: code.Success,
		Msg:  "ok",
		Data: data,
	}
}

func NewResponseFailed() *CommonResp {
	return &CommonResp{
		Code: code.FailedErr.Code(),
		Msg:  code.FailedErr.Msg(),
		Data: nil,
	}
}

func NewResponseFailedBinding(bindErr error) *CommonResp {
	return &CommonResp{
		Code: code.ParamInvalid,
		Msg:  bindErr.Error(),
		Data: nil,
	}
}

func NewResponseFailedMsg(msg string) *CommonResp {
	return &CommonResp{
		Code: code.Failed,
		Msg:  msg,
		Data: nil,
	}
}

func NewResponseFailedCodeMsg(code int, msg string) *CommonResp {
	return &CommonResp{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func NewResponseFailedBizErr(bizErr code.BizErr) *CommonResp {
	return &CommonResp{
		Code: bizErr.Code(),
		Msg:  bizErr.Msg(),
		Data: nil,
	}
}

func NewResponseFailedErr(err error) *CommonResp {
	return &CommonResp{
		Code: code.Failed,
		Msg:  err.Error(),
		Data: nil,
	}
}
