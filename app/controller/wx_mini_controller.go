package controller

import (
	"ginson/app/model"
	"ginson/app/service"
	"ginson/pkg/code"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		if err, ok := err.(validator.ValidationErrors); ok {
			c.FailedWithBizErr(ctx, code.ParamInvalidErr)
		} else {
			c.FailedWithCodeMsg(ctx, code.Failed, err.Error())
		}
		return
	}

	resp, bizErr := c.wxMiniService.WxMiniLogin(ctx, req)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
		return
	}
	c.Success(ctx, resp)
}

func (c *WxMiniController) WxMiniGetUserInfo(ctx *gin.Context) {
	var req *model.WxMiniGetUserInfoReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			c.FailedWithBizErr(ctx, code.ParamInvalidErr)
		} else {
			c.FailedWithCodeMsg(ctx, code.Failed, err.Error())
		}
		return
	}

	resp, bizErr := c.wxMiniService.WxMiniGetUserInfo(ctx, req)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
		return
	}
	c.Success(ctx, resp)
}
