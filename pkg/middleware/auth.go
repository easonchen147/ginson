package middleware

import (
	"ginson/app/service"
	"ginson/pkg/code"
	"ginson/pkg/log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TokenMiddleware Token校验
func TokenMiddleware(userService *service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusOK, code.BizErrorWithCode(code.TokenEmpty))
			return
		}

		userId, err := userService.ParseToken(ctx, strings.TrimSpace(strings.Trim(token, "Bearer")))
		if err != nil || userId == 0 {
			log.Error(ctx, "parse token failed, error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusOK, err)
			return
		}

		log.Info(ctx, "parse token success, userId: %s", userId)
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
