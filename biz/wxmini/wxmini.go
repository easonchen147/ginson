package wxmini

import (
	"context"
	"ginson/biz/user"
	"ginson/foundation/log"
	"ginson/pkg/conf"
	"github.com/easonchen147/thirdpartyoauth"
)

type Service struct {
	wxMiniOauthHandler *thirdpartyoauth.WxMiniOauthHandler
	userService        *user.Service
}

func NewService() *Service {
	return &Service{wxMiniOauthHandler: thirdpartyoauth.NewWxMiniOauthHandler(conf.ExtConf().WxMiniAppId, conf.ExtConf().WxMiniAppSecret), userService: user.NewService()}
}

func (w *Service) WxMiniLogin(ctx context.Context, req *LoginReq) (*user.TokenResp, error) {
	sessionInfo, err := w.wxMiniOauthHandler.CodeToSessionKey(ctx, req.Code)
	if err != nil {
		log.Error(ctx, "code to session key failed, error: %v", err)
		return nil, err
	}

	var userInfo *thirdpartyoauth.WxMiniOauthUserInfo
	if req.EncryptedData != "" && req.Iv != "" {
		userInfo, err = w.wxMiniOauthHandler.GetUserInfo(sessionInfo.SessionKey, req.EncryptedData, req.Iv)
		if err != nil {
			log.Error(ctx, "code to session key failed, error: %v", err)
			return nil, err
		}
	}

	var nickName, avatar string
	var gender int
	if userInfo != nil {
		nickName = userInfo.NickName
		avatar = userInfo.AvatarUrl
		gender = userInfo.Gender
	}

	return w.userService.GetUserToken(ctx, &user.CreateTokenReq{
		OpenId:   sessionInfo.Openid,
		Source:   thirdpartyoauth.SourceWxMini,
		Nickname: nickName,
		Avatar:   avatar,
		Gender:   w.wxMiniOauthHandler.GetGenderByInt(gender),
	})
}
