package controller

import (
	"context"
	"errors"
	"ginson/app/model"
	"ginson/app/service"
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

func (u *UserController) GetUserIdFromCtx(ctx context.Context) (uint, error) {
	userId := ctx.Value("userId")
	if userId == nil {
		return 0, errors.New("userId is nil")
	}

	userIdInt, ok := userId.(uint)
	if !ok {
		return 0, errors.New("userId is invalid")
	}

	return userIdInt, nil
}

func (u *UserController) GetUserInfo(ctx *gin.Context) {
	userId, err := u.GetUserIdFromCtx(ctx)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}

	resp, bizErr := u.userService.GetUserInfo(ctx, userId)
	if bizErr != nil {
		u.FailedWithBizErr(ctx, bizErr)
		return
	}

	u.SuccessData(ctx, resp)
}

func (u *UserController) UpdateUserInfo(ctx *gin.Context) {
	userId, err := u.GetUserIdFromCtx(ctx)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}

	var updateUserInfo *model.UserInfo
	err = ctx.ShouldBindJSON(&updateUserInfo)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}
	updateUserInfo.UserId = userId

	bizErr := u.userService.UpdateUserInfo(ctx, updateUserInfo)
	if bizErr != nil {
		u.FailedWithBizErr(ctx, bizErr)
		return
	}

	u.Success(ctx)
}
