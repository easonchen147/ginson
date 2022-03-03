package middleware

import (
	"fmt"
	"ginson/app/services"
	"ginson/pkg/log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户鉴权
func AuthMiddleware(user *services.UserService) gin.HandlerFunc {
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
			log.Info(fmt.Sprintf("parse token success, userId: %d", userId))
			ctx.Set("userId", userId)
			ctx.Next()
		} else {
			log.Error(fmt.Sprintf("parse token failed, error: %s", err))
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 405,
				"msg":  "用户Token无效",
			})
		}
	}
}
