package user

import (
	"context"
	"errors"
	"ginson/pkg/api"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*api.Handler
	// 放业务使用的service
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		Handler: api.NewHandler(),
		service: NewService(),
	}
}

func (u *Handler) GetUserIdFromCtx(ctx context.Context) (uint, error) {
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

func (u *Handler) GetUserInfo(ctx *gin.Context) {
	userId, err := u.GetUserIdFromCtx(ctx)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}

	resp, bizErr := u.service.GetUserInfo(ctx, userId)
	if bizErr != nil {
		u.FailedWithBizErr(ctx, bizErr)
		return
	}

	u.SuccessData(ctx, resp)
}

func (u *Handler) UpdateUserInfo(ctx *gin.Context) {
	userId, err := u.GetUserIdFromCtx(ctx)
	if err != nil {
		u.FailedWithBindErr(ctx, err)
		return
	}

	var updateUserInfo *UserInfo
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
