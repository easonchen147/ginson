package wxmini

import (
	api2 "ginson/api"
	"ginson/pkg/resp"
	"ginson/service/wxmini"
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

func RegisterWxMiniRouters(group *gin.RouterGroup) {
	wxMini := newHandler()
	group.POST("/login", wxMini.WxMiniLogin)
}

func (h *handler) WxMiniLogin(ctx *gin.Context) {
	var req *api2.LoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		h.FailedWithBindErr(ctx, err)
		return
	}

	var result *api2.TokenResp
	result, err = h.service.WxMiniLogin(ctx, req)
	if err != nil {
		h.FailedWithErr(ctx, err)
		return
	}
	h.SuccessData(ctx, result)
}
