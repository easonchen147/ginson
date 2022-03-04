package router

import (
	"ginson/app/controller"
	"github.com/gin-gonic/gin"
)

type ControllerRoutes interface {
	BindRoutes(group *gin.RouterGroup)
}

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", controller.BaseController.Index)
	r := server.Group("/api")
	{
		userRouter.BindRoutes(r.Group("/user"))
	}
}
