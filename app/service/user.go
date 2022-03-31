package service

import (
	"context"
	"errors"
	"fmt"
	"ginson/app/model"
	"ginson/app/repository/cache"
	"ginson/app/repository/mysql"
	"ginson/pkg/code"
	"ginson/pkg/conf"
	"ginson/pkg/log"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var (
	secret = []byte(conf.AppConf.Ext.TokenSecret)
)

// 定义模块功能错误
var (
	loginFailedErr   = code.BizErrorWithCode(code.LoginFailed)
	tokenInvalidErr  = code.BizErrorWithCode(code.TokenInvalid)
	openIdInvalidErr = code.BizErrorWithCode(code.OpenIdInvalid)
)

type UserService struct {
	userQuery *mysql.UserQuery
	userCache *cache.UserCache
}

var userService = &UserService{userQuery: mysql.GetUserQuery(), userCache: cache.GetUserCache()}

func GetUserService() *UserService {
	return userService
}

func (s *UserService) GetUserInfo(ctx context.Context, userId uint) (*model.UserInfo, code.BizErr) {
	user, err := s.userCache.GetUser(ctx, userId)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, code.BizError(err)
	}

	if errors.Is(err, redis.Nil) {
		user, err = s.queryUserInfoFromDb(ctx, userId)
		if err != nil {
			return nil, code.BizError(err)
		}
		_ = s.userCache.SetUser(ctx, user)
	}
	return user, nil
}

func (s *UserService) queryUserInfoFromDb(ctx context.Context, userId uint) (*model.UserInfo, error) {
	user, err := s.userQuery.GetUserById(ctx, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	result := &model.UserInfo{
		UserId:   user.ID,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Age:      user.Age,
		Gender:   user.Gender,
	}
	return result, nil
}

func (s *UserService) GetUserToken(ctx context.Context, req *model.CreateUserTokenReq) (*model.UserTokenResp, code.BizErr) {
	if req.OpenId == "" {
		return nil, openIdInvalidErr
	}

	var user *model.User
	user, err := s.userQuery.FindByOpenIdAndSource(ctx, req.OpenId, req.Source)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(ctx, "find by openId and source failed. error: %v", err)
		return nil, loginFailedErr
	}

	// 不存在则创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if req.NickName == "" {
			req.NickName = s.randomNickName(ctx)
		}
		user, err = s.createUser(ctx, req)
		if err != nil {
			log.Error(ctx, "create user failed. error: %v", err)
			return nil, loginFailedErr
		}
	}

	token, err := s.createToken(ctx, user.ID)
	if err != nil {
		log.Error(ctx, "create user token failed. error: %v", err)
		return nil, loginFailedErr
	}

	resp := &model.UserTokenResp{
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Token:    token,
	}
	return resp, nil
}

func (s *UserService) randomNickName(ctx context.Context) string {
	return fmt.Sprintf("用户%06d", rand.Intn(1000000))
}

func (s *UserService) createUser(ctx context.Context, req *model.CreateUserTokenReq) (*model.User, error) {
	user := &model.User{
		OpenId:   req.OpenId,
		Source:   req.Source,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Age:      req.Age,
		Gender:   req.Gender,
	}
	err := s.userQuery.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.userQuery.FindByOpenIdAndSource(ctx, req.OpenId, req.Source)
}

// ParseToken 解析token
func (s *UserService) ParseToken(ctx context.Context, tokenString string) (int, code.BizErr) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, tokenInvalidErr
		}
		return secret, nil
	})
	if err != nil {
		return 0, tokenInvalidErr
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["userId"].(float64)), nil
	}

	return 0, tokenInvalidErr
}

// 创建token
func (s *UserService) createToken(ctx context.Context, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
