package user

import (
	"context"
	"errors"
	"fmt"
	"ginson/biz/user/db"
	"ginson/foundation/log"
	"ginson/foundation/util"
	"ginson/pkg/code"
	"ginson/pkg/constant"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	query *db.Query
	cache *Cache
}

func NewService() *Service {
	return &Service{query: db.NewQuery(), cache: NewCache()}
}

func (u *Service) GetUserInfo(ctx context.Context, userId uint) (*User, error) {
	user, err := u.cache.GetUser(ctx, userId)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, code.RedisError
	}

	if errors.Is(err, redis.Nil) {
		user, err = u.queryUserInfoFromDb(ctx, userId)
		if err != nil {
			return nil, code.MysqlError
		}

		cpyCtx := util.CopyCtx(ctx)
		_ = util.GoInPool(func() {
			_ = u.cache.SetUser(cpyCtx, user)
		})
	}
	return user, nil
}

func (u *Service) UpdateUserInfo(ctx context.Context, userInfo *User) error {
	err := u.query.UpdateUserById(ctx, &db.User{
		Model: gorm.Model{
			ID: userInfo.UserId,
		},
		Nickname: userInfo.Nickname,
		Avatar:   userInfo.Avatar,
		Age:      userInfo.Age,
		Gender:   userInfo.Gender,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *Service) queryUserInfoFromDb(ctx context.Context, userId uint) (*User, error) {
	user, err := u.query.GetUserById(ctx, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	result := &User{
		UserId:   user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Age:      user.Age,
		Gender:   user.Gender,
	}
	return result, nil
}

func (u *Service) GetUserToken(ctx context.Context, req *CreateTokenReq) (*TokenResp, error) {
	if req.OpenId == "" {
		return nil, code.OpenIdInvalidError
	}

	var user *db.User
	user, err := u.query.FindByOpenIdAndSource(ctx, req.OpenId, req.Source)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(ctx, "find by openId and source failed. error: %v", err)
		return nil, code.MysqlError
	}

	// 不存在则创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if req.Nickname == "" {
			req.Nickname = u.randomNickName(ctx)
		}
		user, err = u.createUser(ctx, req)
		if err != nil {
			log.Error(ctx, "create user failed. error: %v", err)
			return nil, code.MysqlError
		}
	}

	if user == nil {
		log.Error(ctx, "user is nil")
		return nil, code.LoginFailedError
	}

	token, err := u.createToken(ctx, user.ID)
	if err != nil {
		log.Error(ctx, "create user token failed. error: %v", err)
		return nil, err
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

func (u *Service) createUser(ctx context.Context, req *CreateTokenReq) (*db.User, error) {
	user := &db.User{
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
