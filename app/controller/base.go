package controller

import (
	"ginson/app/model/resp"
	"ginson/pkg/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

var BaseController = &Controller{}

func (*Controller) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to ginson")
}

func (*Controller) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, resp.NewResponseSuccess(nil))
}

func (*Controller) SuccessData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, resp.NewResponseSuccess(data))
}

func (*Controller) Failed(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailed())
}

func (*Controller) FailedWithBindErr(ctx *gin.Context, bindErr error) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailedBinding(bindErr))
}

func (*Controller) FailedWithMsg(ctx *gin.Context, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailedMsg(msg))
}

func (*Controller) FailedWithCodeMsg(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailedCodeMsg(code, msg))
}

func (*Controller) FailedWithBizErr(ctx *gin.Context, bizErr code.BizErr) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailedBizErr(bizErr))
}

func (*Controller) FailedWithErr(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailedErr(err))
}
