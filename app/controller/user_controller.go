package controller

import (
	"ginson/app/model"
	"ginson/app/service"
	"ginson/pkg/log"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	*Controller
	// 放业务使用的service
	userService *service.UserService
}

var userController = &UserController{
	Controller:  BaseController,
	userService: service.GetUserService(),
}

func GetUserController() *UserController {
	return userController
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var form model.UserRegisterCommand
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		c.FailedWithBindErr(ctx, err)
		return
	}
	token, bizErr := c.userService.Register(ctx, &form)
	if bizErr != nil {
		c.FailedWithBizErr(ctx, bizErr)
	} else {
		c.Success(ctx, gin.H{"token": token})
	}
	return
}

// Login 用户登录
func (c *UserController) Login(ctx *gin.Context) {
	var form model.UserLoginCommand
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		c.FailedWithBindErr(ctx, err)
		return
	}
	token, bizErr := c.userService.Login(ctx, &form)
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

	log.Debug("add token into blacklist, token: %s", token)

	c.Success(ctx, "")
}
