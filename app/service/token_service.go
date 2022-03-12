package service

import (
	"context"
	"fmt"
	"ginson/pkg/code"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenService struct{}

var tokenService = &TokenService{}

func GetTokenService() *TokenService {
	return tokenService
}

// ParseToken 解析token
func (s *TokenService) ParseToken(ctx context.Context, tokenString string) (string, code.BizErr) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return "", code.BizError(err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["openId"].(string), nil
	} else {
		return "", code.BizError(err)
	}
}

// 创建token
func (s *TokenService) createToken(ctx context.Context, openId, source string) (string, code.BizErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"openId": openId,
		"source": source,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", code.BizError(err)
	}
	return tokenString, nil
}
