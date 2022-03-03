package controllers

import (
	"fmt"
	"ginson/app/models"
	"ginson/app/services"
	"ginson/pkg/code"
	"ginson/pkg/log"
	"ginson/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	*Controller
	*services.UserService
}

func (c *UserController) BindRoutes(group *gin.RouterGroup) {
	group.POST("/register", userController.Register) //用户注册
	group.POST("/login", userController.Login)       //用户登录
	group.Use(middleware.AuthMiddleware(c.UserService)).
		POST("/logout", userController.Logout) //用户退出登录
}

var userController = &UserController{
	Controller:  BaseController,
	UserService: services.GetUserService(),
}

func GetUserController() *UserController {
	return userController
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var form models.UserRegisterCommand
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			c.FailedWithBizErr(ctx, code.ParamInvalidErr)
		} else {
			c.FailedWithCodeMsg(ctx, code.Failed, err.Error())
		}
		return
	}
	token, bizErr := c.UserService.Register(ctx, &form)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
	} else {
		c.Success(ctx, gin.H{"token": token})
	}
	return
}

// Login 用户登录
func (c *UserController) Login(ctx *gin.Context) {
	var form models.UserLoginCommand
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			c.FailedWithBizErr(ctx, code.ParamInvalidErr)
		} else {
			c.FailedWithCodeMsg(ctx, code.Failed, err.Error())
		}
		return
	}
	token, bizErr := c.UserService.Login(ctx, &form)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
	} else {
		c.Success(ctx, gin.H{"token": token})
	}
	return
}

// Logout 用户退出
func (c *UserController) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	log.Debug(fmt.Sprintf("add token into blacklist, token: %s", token))

	c.Success(ctx, "")
}
