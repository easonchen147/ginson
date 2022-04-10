package user

import (
	"context"
	"errors"
	"ginson/pkg/resp"

	"github.com/gin-gonic/gin"
)

type handler struct {
	*resp.Handler
	// 放业务使用的service
	service *Service
}

func newHandler() *handler {
	return &handler{
		Handler: resp.NewHandler(),
		service: NewService(),
	}
}

func (u *handler) GetUserIdFromCtx(ctx context.Context) (uint, error) {
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

func (u *handler) GetUserInfo(ctx *gin.Context) {
	userId, err := u.GetUserIdFromCtx(ctx)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}

	result, bizErr := u.service.GetUserInfo(ctx, userId)
	if bizErr != nil {
		u.FailedWithBizErr(ctx, bizErr)
		return
	}

	u.SuccessData(ctx, result)
}

func (u *handler) UpdateUserInfo(ctx *gin.Context) {
	userId, err := u.GetUserIdFromCtx(ctx)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}

	var updateUserInfo *Info
	err = ctx.ShouldBindJSON(&updateUserInfo)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}
	updateUserInfo.UserId = userId

	bizErr := u.service.UpdateUserInfo(ctx, updateUserInfo)
	if bizErr != nil {
		u.FailedWithBizErr(ctx, bizErr)
		return
	}

	u.Success(ctx)
}
