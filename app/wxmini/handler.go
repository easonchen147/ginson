package wxmini

import (
	"ginson/pkg/api"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*api.Handler
	// 放业务使用的service
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		Handler: api.NewHandler(),
		service: NewService(),
	}
}

func (c *Handler) WxMiniLogin(ctx *gin.Context) {
	var req *WxMiniLoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		c.FailedWithBindErr(ctx, err)
		return
	}

	resp, bizErr := c.service.WxMiniLogin(ctx, req)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
		return
	}
	c.SuccessData(ctx, resp)
}
