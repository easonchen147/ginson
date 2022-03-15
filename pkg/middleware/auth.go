package middleware

import (
	"ginson/app/service"
	"ginson/pkg/log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户鉴权
func AuthMiddleware(user *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 405,
				"msg":  "未登录",
			})
			return
		}
		userId, err := user.ParseToken(ctx, strings.TrimSpace(strings.Trim(token, "Bearer")))
		if err == nil && userId > 0 {
			log.Info("parse token success, userId: %d", userId)
			ctx.Set("userId", userId)
			ctx.Next()
		} else {
			log.Error("parse token failed, error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 405,
				"msg":  "用户Token无效",
			})
		}
	}
}

// TokenMiddleware Token校验
func TokenMiddleware(tokenService *service.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 405,
				"msg":  "未登录",
			})
			return
		}
		openId, err := tokenService.ParseToken(ctx, strings.TrimSpace(strings.Trim(token, "Bearer")))
		if err == nil && openId != "" {
			log.Info("parse token success, openId: %s", openId)
			ctx.Set("openId", openId)
			ctx.Next()
		} else {
			log.Error("parse token failed, error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 405,
				"msg":  "用户Token无效",
			})
		}
	}
}
