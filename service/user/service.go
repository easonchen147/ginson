package user

import (
	"context"
	"errors"
	"fmt"
	"ginson/api"
	"ginson/model"
	"ginson/repo/user/cache"
	"ginson/repo/user/db"
	"math/rand"
	"time"

	"ginson/pkg/code"
	"ginson/pkg/conf"

	"github.com/easonchen147/foundation/log"
	"github.com/easonchen147/foundation/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service struct {
	db    db.Repodb
	cache cache.RepoCache
}

func NewService() *Service {
	return &Service{db: db.NewDb(), cache: cache.NewRds()}
}

func (u *Service) GetUserInfo(ctx context.Context, userId uint) (*api.UserVO, error) {
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

func (u *Service) UpdateUserInfo(ctx context.Context, userInfo *api.UserVO) error {
	err := u.db.UpdateUserById(ctx, &model.User{
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

func (u *Service) queryUserInfoFromDb(ctx context.Context, userId uint) (*api.UserVO, error) {
	user, err := u.db.GetUserById(ctx, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	result := &api.UserVO{
		UserId:   user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Age:      user.Age,
		Gender:   user.Gender,
	}
	return result, nil
}

func (u *Service) GetUserToken(ctx context.Context, req *api.CreateTokenReq) (*api.TokenResp, error) {
	if req.OpenId == "" {
		return nil, code.OpenIdInvalidError
	}

	var user *model.User
	user, err := u.db.FindByOpenIdAndSource(ctx, req.OpenId, req.Source)
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

	resp := &api.TokenResp{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Token:    token,
	}
	return resp, nil
}

func (u *Service) randomNickName(ctx context.Context) string {
	return fmt.Sprintf("用户%06d", rand.Intn(1000000))
}

func (u *Service) createUser(ctx context.Context, req *api.CreateTokenReq) (*model.User, error) {
	user := &model.User{
		OpenId:   req.OpenId,
		Source:   req.Source,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Age:      req.Age,
		Gender:   req.Gender,
	}
	err := u.db.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

var tokenSecret = []byte(conf.ExtConf().TokenSecret)

// 创建token
func (u *Service) createToken(ctx context.Context, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
