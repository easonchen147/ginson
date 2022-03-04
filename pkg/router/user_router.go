package router

import (
	"ginson/app/controller"
	"ginson/app/service"
	"ginson/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	controller *controller.UserController
}

var userRouter = &UserRouter{controller: controller.GetUserController()}

func (c *UserRouter) BindRoutes(group *gin.RouterGroup) {
	group.POST("/register", c.controller.Register) //用户注册
	group.POST("/login", c.controller.Login)       //用户登录
	group.Use(middleware.AuthMiddleware(service.GetUserService())).
		POST("/logout", c.controller.Logout) //用户退出登录
}
