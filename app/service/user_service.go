package service

import (
	"context"
	"errors"
	"fmt"
	"ginson/app/model"
	"ginson/app/repository/mysql"
	"ginson/pkg/code"
	"ginson/pkg/conf"
	"ginson/pkg/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var (
	secret = []byte(conf.AppConf.Ext.TokenSecret)
)

// 定义模块功能错误
var (
	emailExistsErr     = code.BizErrorWithCode(code.EmailExists)
	emailNotExistsErr  = code.BizErrorWithCode(code.EmailNotExists)
	emailOrPswWrongErr = code.BizErrorWithCode(code.EmailOrPswWrong)
	loginFailedErr     = code.BizErrorWithCode(code.LoginFailed)
)

type UserService struct {
	userQuery *mysql.UserQuery
}

var userService = &UserService{userQuery: mysql.GetUserQuery()}

func GetUserService() *UserService {
	return userService
}

func (s *UserService) Register(ctx context.Context, param *model.UserRegisterCommand) (string, code.BizErr) {
	emailExisted, err := s.userQuery.IsUserEmailExisted(ctx, param.Email)
	if err != nil {
		return "", code.FailedErr
	}
	if emailExisted {
		return "", emailExistsErr
	}
	var user = model.User{}
	salt := utils.GetUuidV4()[24:]
	user.Name = param.Name
	user.Email = param.Email
	user.Password = utils.Sha1([]byte(param.Password + salt))
	user.Salt = salt
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = s.userQuery.CreateUser(ctx, &user)
	if err != nil {
		return "", code.BizError(err)
	} else {
		token, err := s.createToken(ctx, user.Id)
		if err != nil {
			return "", code.BizError(err)
		} else {
			return token, nil
		}
	}
}

func (s *UserService) Login(ctx context.Context, param *model.UserLoginCommand) (string, code.BizErr) {
	var user *model.User
	user, err := s.userQuery.FindUserByEmail(ctx, param.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", emailNotExistsErr
		}
		return "", code.BizError(err)
	}
	if user.Password != utils.Sha1([]byte(param.Password+user.Salt)) {
		return "", emailOrPswWrongErr
	} else {
		token, err := s.createToken(ctx, user.Id)
		if err != nil {
			return "", loginFailedErr
		} else {
			return token, nil
		}
	}
}

// ParseToken 解析token
func (s *UserService) ParseToken(ctx context.Context, tokenString string) (int, code.BizErr) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return 0, code.BizError(err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["userId"].(float64)), nil
	} else {
		return 0, code.BizError(err)
	}
}

// 创建token
func (s *UserService) createToken(ctx context.Context, userId uint) (string, code.BizErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", code.BizError(err)
	}
	return tokenString, nil
}
