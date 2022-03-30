package router

import (
	"ginson/app/controller"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	controller *controller.UserController
}

var userRouter = &UserRouter{controller: controller.GetUserController()}

func (c *UserRouter) BindRoutes(group *gin.RouterGroup) {
	group.POST("/get-user-info", c.controller.GetUserInfo)
}
