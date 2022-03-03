package controllers

import (
	"ginson/pkg/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerRoutes interface {
	BindRoutes(group *gin.RouterGroup)
}

type Controller struct{}

type commonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var BaseController = &Controller{}

func (*Controller) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Gen Web")
}

func (*Controller) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, commonResp{
		Code: code.Success,
		Msg:  "ok",
		Data: data,
	})
}

func (*Controller) Failed(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, commonResp{
		Code: code.FailedErr.Code(),
		Msg:  code.FailedErr.Msg(),
		Data: nil,
	})
}

func (*Controller) FailedWithCodeMsg(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, commonResp{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func (*Controller) FailedWithBizErr(ctx *gin.Context, bizErr code.BizErr) {
	ctx.AbortWithStatusJSON(http.StatusOK, commonResp{
		Code: bizErr.Code(),
		Msg:  bizErr.Msg(),
		Data: nil,
	})
}

func (*Controller) FailedWithErr(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusOK, commonResp{
		Code: code.Failed,
		Msg:  err.Error(),
		Data: nil,
	})
}
