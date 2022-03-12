package router

import (
	"ginson/app/controller"
	"github.com/gin-gonic/gin"
)

type WxMiniRouter struct {
	controller *controller.WxMiniController
}

var wxMiniRouter = &WxMiniRouter{controller: controller.GetWxMiniController()}

func (c *WxMiniRouter) BindRoutes(group *gin.RouterGroup) {
	group.POST("/login", c.controller.WxMiniLogin)               //微信授权
	group.POST("/get-user-info", c.controller.WxMiniGetUserInfo) //微信用户信息获取
}
