package user

import (
	"context"
	"errors"
	"ginson/api"
	"ginson/pkg/middleware"
	"ginson/pkg/resp"
	"ginson/service/user"
	"github.com/gin-gonic/gin"
)

type handler struct {
	*resp.Handler
	// 放业务使用的service
	service *user.Service
}

func newHandler() *handler {
	return &handler{
		Handler: resp.NewHandler(),
		service: user.NewService(),
	}
}

func RegisterUserRouters(group *gin.RouterGroup) {
	userHandler := newHandler()
	group.Use(middleware.TokenMiddleware())
	group.POST("/get-user-info", userHandler.GetUserInfo)
}

func (h *handler) GetUserIdFromCtx(ctx context.Context) (uint, error) {
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

func (h *handler) GetUserInfo(ctx *gin.Context) {
	userId, err := h.GetUserIdFromCtx(ctx)
	if err != nil {
		h.FailedWithBindErr(ctx, err)
		return
	}

	var result *api.UserVO
	result, err = h.service.GetUserInfo(ctx, userId)
	if err != nil {
		h.FailedWithErr(ctx, err)
		return
	}

	h.SuccessData(ctx, result)
}

func (h *handler) UpdateUserInfo(ctx *gin.Context) {
	userId, err := h.GetUserIdFromCtx(ctx)
	if err != nil {
		h.FailedWithErr(ctx, err)
		return
	}

	var updateUserInfo *api.UserVO
	err = ctx.ShouldBindJSON(&updateUserInfo)
	if err != nil {
		h.FailedWithBindErr(ctx, err)
		return
	}
	updateUserInfo.UserId = userId

	err = h.service.UpdateUserInfo(ctx, updateUserInfo)
	if err != nil {
		h.FailedWithErr(ctx, err)
		return
	}

	h.Success(ctx)
}
