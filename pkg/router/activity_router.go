package router

import (
	"ginson/app/controller"
	"ginson/app/service"
	"ginson/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct {
	controller *controller.ActivityController
}

var activityRouter = &ActivityRouter{controller: controller.GetActivityController()}

func (c *ActivityRouter) BindRoutes(group *gin.RouterGroup) {
	//group.Use(middleware.TokenMiddleware(service.GetTokenService()))
	group.POST("/get-prize", c.controller.GetPrize, middleware.TokenMiddleware(service.GetTokenService()))
	group.GET("/get-qr-code", c.controller.GetQrCode)
	group.GET("/get-screenshot", c.controller.GetScreenShot)
}
