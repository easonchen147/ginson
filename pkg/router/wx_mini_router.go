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
	group.POST("/login", c.controller.WxMiniLogin)
}
