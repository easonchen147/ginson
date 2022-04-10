package wxmini

import (
	"ginson/pkg/resp"

	"github.com/gin-gonic/gin"
)

type handler struct {
	*resp.Handler
	// 放业务使用的service
	service *Service
}

func newHandler() *handler {
	return &handler{
		Handler: resp.NewHandler(),
		service: NewService(),
	}
}

func (c *handler) WxMiniLogin(ctx *gin.Context) {
	var req *LoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		c.FailedWithBindErr(ctx, err)
		return
	}

	result, bizErr := c.service.WxMiniLogin(ctx, req)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
		return
	}
	c.SuccessData(ctx, result)
}
