package router

import (
	"ginson/app/controller"
	"github.com/gin-gonic/gin"
)

type ToolRouter struct {
	controller *controller.ToolController
}

var toolRouter = &ToolRouter{controller: controller.GetToolController()}

func (c *ToolRouter) BindRoutes(group *gin.RouterGroup) {
	group.GET("/get-qr-code", c.controller.GetQrCode)
	group.GET("/get-screenshot", c.controller.GetScreenShot)
}
