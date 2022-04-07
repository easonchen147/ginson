package wxmini

import (
	"context"
	"ginson/app/service/user"
	"ginson/conf"
	"ginson/pkg/code"
	"ginson/pkg/constant"
	"ginson/pkg/log"
	"ginson/pkg/oauth"
)

type Service struct {
	wxMiniOauthHandler *oauth.WxMiniOauthHandler
	userService        *user.Service
}

func NewService() *Service {
	return &Service{wxMiniOauthHandler: oauth.NewWxMiniOauthHandler(conf.AppConf.Ext.WxMiniAppId, conf.AppConf.Ext.WxMiniAppSecret), userService: user.NewService()}
}

func (w *Service) WxMiniLogin(ctx context.Context, req *WxMiniLoginReq) (*user.UserTokenResp, code.BizErr) {
	sessionInfo, err := w.wxMiniOauthHandler.CodeToSessionKey(ctx, req.Code)
	if err != nil {
		log.Error(ctx, "code to session key failed, error: %v", err)
		return nil, code.BizError(err)
	}

	var userInfo *oauth.WxMiniOauthUserInfo
	if req.EncryptedData != "" && req.Iv != "" {
		userInfo, err = w.wxMiniOauthHandler.GetUserInfo(sessionInfo.SessionKey, req.EncryptedData, req.Iv)
		if err != nil {
			log.Error(ctx, "code to session key failed, error: %v", err)
			return nil, code.BizError(err)
		}
	}

	var nickName, avatar string
	var gender int
	if userInfo != nil {
		nickName = userInfo.NickName
		avatar = userInfo.AvatarUrl
		gender = userInfo.Gender
	}

	return w.userService.GetUserToken(ctx, &user.CreateUserTokenReq{
		OpenId:   sessionInfo.Openid,
		Source:   constant.OauthSourceWxMini,
		Nickname: nickName,
		Avatar:   avatar,
		Gender:   w.wxMiniOauthHandler.GetGenderByInt(gender),
	})
}
