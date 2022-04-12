package resp

import (
	"ginson/pkg/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to ginson")
}

func (*Handler) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewResponseSuccess(nil))
}

func (*Handler) SuccessData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, NewResponseSuccess(data))
}

func (*Handler) Failed(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, NewResponseFailed())
}

func (*Handler) FailedWithBindErr(ctx *gin.Context, bindErr error) {
	ctx.AbortWithStatusJSON(http.StatusOK, NewResponseFailedBinding(bindErr))
}

func (*Handler) FailedWithMsg(ctx *gin.Context, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, NewResponseFailedMsg(msg))
}

func (*Handler) FailedWithCodeMsg(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, NewResponseFailedCodeMsg(code, msg))
}

func (*Handler) FailedWithErr(ctx *gin.Context, err error) {
	if bizErr, ok := err.(code.BizErr); ok {
		ctx.AbortWithStatusJSON(http.StatusOK, NewResponseFailedBizErr(bizErr))
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, NewResponseFailedErr(err))
}
