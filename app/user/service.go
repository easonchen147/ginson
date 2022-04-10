package user

import (
	"context"
	"errors"
	"fmt"
	"ginson/pkg/code"
	"ginson/pkg/constant"
	"ginson/pkg/log"
	"ginson/pkg/util"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Service struct {
	query *query
	cache *cache
}

func NewService() *Service {
	return &Service{query: newQuery(), cache: newCache()}
}

func (u *Service) GetUserInfo(ctx context.Context, userId uint) (*Info, code.BizErr) {
	user, err := u.cache.GetUser(ctx, userId)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, code.BizError(err)
	}

	if errors.Is(err, redis.Nil) {
		user, err = u.queryUserInfoFromDb(ctx, userId)
		if err != nil {
			return nil, code.BizError(err)
		}

		cpyCtx := util.CopyCtx(ctx)
		_ = util.GoInPool(func() {
			_ = u.cache.SetUser(cpyCtx, user)
		})
	}
	return user, nil
}

func (u *Service) UpdateUserInfo(ctx context.Context, userInfo *Info) code.BizErr {
	err := u.query.UpdateUserById(ctx, userInfo)
	if err != nil {
		return code.BizError(err)
	}
	return nil
}

func (u *Service) queryUserInfoFromDb(ctx context.Context, userId uint) (*Info, error) {
	user, err := u.query.GetUserById(ctx, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	result := &Info{
		UserId:   user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Age:      user.Age,
		Gender:   user.Gender,
	}
	return result, nil
}

func (u *Service) GetUserToken(ctx context.Context, req *CreateTokenReq) (*TokenResp, code.BizErr) {
	if req.OpenId == "" {
		return nil, code.OpenIdInvalidErr
	}

	var user *User
	user, err := u.query.FindByOpenIdAndSource(ctx, req.OpenId, req.Source)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(ctx, "find by openId and source failed. error: %v", err)
		return nil, code.LoginFailedErr
	}

	// 不存在则创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if req.Nickname == "" {
			req.Nickname = u.randomNickName(ctx)
		}
		user, err = u.createUser(ctx, req)
		if err != nil {
			log.Error(ctx, "create user failed. error: %v", err)
			return nil, code.LoginFailedErr
		}
	}

	if user == nil {
		log.Error(ctx, "user is nil")
		return nil, code.ServerErr
	}

	token, err := u.createToken(ctx, user.ID)
	if err != nil {
		log.Error(ctx, "create user token failed. error: %v", err)
		return nil, code.LoginFailedErr
	}

	resp := &TokenResp{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Token:    token,
	}
	return resp, nil
}

func (u *Service) randomNickName(ctx context.Context) string {
	return fmt.Sprintf("用户%06d", rand.Intn(1000000))
}

func (u *Service) createUser(ctx context.Context, req *CreateTokenReq) (*User, error) {
	user := &User{
		OpenId:   req.OpenId,
		Source:   req.Source,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Age:      req.Age,
		Gender:   req.Gender,
	}
	err := u.query.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 创建token
func (u *Service) createToken(ctx context.Context, userId uint) (string, error) {
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
