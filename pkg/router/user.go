package router

import (
	"ginson/app/controller"
	"ginson/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	controller *controller.UserController
}

var userRouter = &UserRouter{controller: controller.GetUserController()}

func (c *UserRouter) BindRoutes(group *gin.RouterGroup) {
	group.Use(middleware.TokenMiddleware())
	group.POST("/get-user-info", c.controller.GetUserInfo)
}
