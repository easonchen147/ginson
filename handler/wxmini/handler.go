package wxmini

import (
	"ginson/api"
	"ginson/pkg/resp"
	"ginson/service/wxmini"
	wx "ginson/wxmini"

	"github.com/gin-gonic/gin"
)

type handler struct {
	*resp.Handler
	// 放业务使用的service
	service *wxmini.Service
}

func newHandler() *handler {
	return &handler{
		Handler: resp.NewHandler(),
		service: wxmini.NewService(),
	}
}

func (c *handler) WxMiniLogin(ctx *gin.Context) {
	var req *wx.LoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		c.FailedWithBindErr(ctx, err)
		return
	}

	var result *api.TokenResp
	result, err = c.service.WxMiniLogin(ctx, req)
	if err != nil {
		c.FailedWithErr(ctx, err)
		return
	}
	c.SuccessData(ctx, result)
}
