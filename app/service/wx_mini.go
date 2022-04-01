package service

import (
	"context"
	"ginson/app/model"
	"ginson/pkg/code"
	"ginson/pkg/conf"
	"ginson/pkg/constant"
	"ginson/pkg/log"
	"ginson/pkg/oauth"
)

var wxMiniOauthHandler = oauth.NewWxMiniOauthHandler(conf.AppConf.Ext.WxMiniAppId, conf.AppConf.Ext.WxMiniAppSecret)

type WxMiniService struct {
	wxMiniOauthHandler *oauth.WxMiniOauthHandler
}

var wxMiniService = &WxMiniService{wxMiniOauthHandler: wxMiniOauthHandler}

func GetWxMiniService() *WxMiniService {
	return wxMiniService
}

func (w *WxMiniService) WxMiniLogin(ctx context.Context, req *model.WxMiniLoginReq) (*model.UserTokenResp, code.BizErr) {
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

	return userService.GetUserToken(ctx, &model.CreateUserTokenReq{
		OpenId:   sessionInfo.Openid,
		Source:   constant.OauthSourceWxMini,
		Nickname: nickName,
		Avatar:   avatar,
		Gender:   w.wxMiniOauthHandler.GetGenderByInt(gender),
	})
}
