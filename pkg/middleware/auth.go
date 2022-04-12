package middleware

import (
	"context"
	"ginson/pkg/code"
	"ginson/pkg/constant"
	"ginson/pkg/log"
	"ginson/pkg/resp"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenMiddleware Token校验
func TokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailedBizErr(code.TokenEmptyError))
			return
		}

		userId, err := parseToken(ctx, token)
		if err != nil || userId == 0 {
			log.Error(ctx, "parse token failed, error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusOK, resp.NewResponseFailedBizErr(err))
			return
		}

		log.Info(ctx, "parse token success, userId: %d", userId)
		ctx.Set("userId", userId)
		ctx.Next()
	}
}

// 解析token
func parseToken(ctx context.Context, tokenString string) (uint, code.BizErr) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, code.TokenInvalidError
		}
		return constant.TokenSecret, nil
	})
	if err != nil {
		return 0, code.TokenInvalidError
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["userId"].(float64)), nil
	}

	return 0, code.TokenInvalidError
}
