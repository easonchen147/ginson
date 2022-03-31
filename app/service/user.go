package service

import (
	"context"
	"errors"
	"fmt"
	"ginson/app/model"
	"ginson/app/repository/cache"
	"ginson/app/repository/mysql"
	"ginson/pkg/code"
	"ginson/pkg/constant"
	"ginson/pkg/log"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type UserService struct {
	db    *mysql.UserDb
	cache *cache.UserCache
}

var userService = &UserService{db: mysql.GetUserDb(), cache: cache.GetUserCache()}

func GetUserService() *UserService {
	return userService
}

func (u *UserService) GetUserInfo(ctx context.Context, userId uint) (*model.UserInfo, code.BizErr) {
	user, err := u.cache.GetUser(ctx, userId)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, code.BizError(err)
	}

	if errors.Is(err, redis.Nil) {
		user, err = u.queryUserInfoFromDb(ctx, userId)
		if err != nil {
			return nil, code.BizError(err)
		}
		_ = u.cache.SetUser(ctx, user)
	}
	return user, nil
}

func (u *UserService) queryUserInfoFromDb(ctx context.Context, userId uint) (*model.UserInfo, error) {
	user, err := u.db.GetUserById(ctx, userId)
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

func (u *UserService) GetUserToken(ctx context.Context, req *model.CreateUserTokenReq) (*model.UserTokenResp, code.BizErr) {
	if req.OpenId == "" {
		return nil, code.OpenIdInvalidErr
	}

	var user *model.User
	user, err := u.db.FindByOpenIdAndSource(ctx, req.OpenId, req.Source)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(ctx, "find by openId and source failed. error: %v", err)
		return nil, code.LoginFailedErr
	}

	// 不存在则创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if req.NickName == "" {
			req.NickName = u.randomNickName(ctx)
		}
		user, err = u.createUser(ctx, req)
		if err != nil {
			log.Error(ctx, "create user failed. error: %v", err)
			return nil, code.LoginFailedErr
		}
	}

	token, err := u.createToken(ctx, user.ID)
	if err != nil {
		log.Error(ctx, "create user token failed. error: %v", err)
		return nil, code.LoginFailedErr
	}

	resp := &model.UserTokenResp{
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Token:    token,
	}
	return resp, nil
}

func (u *UserService) randomNickName(ctx context.Context) string {
	return fmt.Sprintf("用户%06d", rand.Intn(1000000))
}

func (u *UserService) createUser(ctx context.Context, req *model.CreateUserTokenReq) (*model.User, error) {
	user := &model.User{
		OpenId:   req.OpenId,
		Source:   req.Source,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Age:      req.Age,
		Gender:   req.Gender,
	}
	err := u.db.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return u.db.FindByOpenIdAndSource(ctx, req.OpenId, req.Source)
}

// 创建token
func (u *UserService) createToken(ctx context.Context, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(constant.TokenSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
