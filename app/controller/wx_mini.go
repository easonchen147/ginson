package controller

import (
	"ginson/app/model"
	"ginson/app/service"
	"github.com/gin-gonic/gin"
)

type WxMiniController struct {
	*Controller
	// 放业务使用的service
	wxMiniService *service.WxMiniService
}

var wxMiniController = &WxMiniController{
	Controller:    BaseController,
	wxMiniService: service.GetWxMiniService(),
}

func GetWxMiniController() *WxMiniController {
	return wxMiniController
}

func (c *WxMiniController) WxMiniLogin(ctx *gin.Context) {
	var req *model.WxMiniLoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		c.FailedWithBindErr(ctx, err)
		return
	}

	resp, bizErr := c.wxMiniService.WxMiniLogin(ctx, req)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
		return
	}
	c.SuccessData(ctx, resp)
}
