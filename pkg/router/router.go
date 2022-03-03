package router

import (
	"ginson/app/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", controllers.BaseController.Index)
	r := server.Group("/api")
	{
		controllers.GetUserController().BindRoutes(r.Group("/user"))
	}
}
