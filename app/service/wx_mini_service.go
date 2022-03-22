package service

import (
	"context"
	"errors"
	"ginson/app/model"
	"ginson/app/repository/cache"
	"ginson/pkg/code"
	"ginson/pkg/conf"
	"ginson/pkg/constant"
	"ginson/pkg/log"
	"ginson/pkg/oauth"
	"ginson/pkg/utils"
)

var wxMiniOauthHandler = oauth.NewWxMiniOauthHandler(conf.AppConf.Ext.WxMiniAppId, conf.AppConf.Ext.WxMiniAppSecret)

type WxMiniService struct {
	wxMiniOauthHandler *oauth.WxMiniOauthHandler
	wxMiniCache        *cache.WxMiniCache
}

var wxMiniService = &WxMiniService{wxMiniOauthHandler: wxMiniOauthHandler, wxMiniCache: cache.GetWxMiniCache()}

func GetWxMiniService() *WxMiniService {
	return wxMiniService
}

func (w *WxMiniService) WxMiniLogin(ctx context.Context, req *model.WxMiniLoginReq) (*model.WxMiniLoginResp, code.BizErr) {
	sessionInfo, err := w.wxMiniOauthHandler.CodeToSessionKey(ctx, req.Code)
	if err != nil {
		log.Error(ctx, "code to session key failed, error: %v", err)
		return nil, code.BizError(err)
	}

	token, err := tokenService.createToken(ctx, sessionInfo.Openid, constant.OauthSourceWxMini)
	if err != nil {
		return nil, code.BizError(err)
	}

	result := &model.WxMiniLoginResp{Token: token}
	if req.EncryptedData != "" && req.Iv != "" {
		userInfo, err := w.wxMiniOauthHandler.GetUserInfo(sessionInfo.SessionKey, req.EncryptedData, req.Iv)
		if err != nil {
			log.Error(ctx, "code to session key failed, error: %v", err)
			return nil, code.BizError(err)
		}
		result.UserInfo = w.populateUserInfoResp(userInfo)
	}

	_ = utils.GoInPool(func() {
		_ = w.wxMiniCache.AddSessionKey(ctx, sessionInfo.Openid, sessionInfo.SessionKey)
	})

	return result, nil
}

func (w *WxMiniService) WxMiniGetUserInfo(ctx context.Context, req *model.WxMiniGetUserInfoReq) (*model.WxMiniGetUserInfoResp, code.BizErr) {
	openId, ok := ctx.Value("openId").(string)
	if !ok {
		return nil, code.BizError(errors.New("get open id failed"))
	}

	sessionKey, err := w.wxMiniCache.GetSessionKey(ctx, openId)
	if err != nil || sessionKey == "" {
		return nil, code.BizError(errors.New("sessionKey is invalid"))
	}

	userInfo, err := w.wxMiniOauthHandler.GetUserInfo(sessionKey, req.EncryptedData, req.Iv)
	if err != nil {
		log.Error(ctx, "code to session key failed, error: %v", err)
		return nil, code.BizError(err)
	}

	result := w.populateUserInfoResp(userInfo)

	_ = utils.GoInPool(func() {
		_ = w.wxMiniCache.AddUserInfo(ctx, openId, result)
	})

	return result, nil
}

func (w *WxMiniService) populateUserInfoResp(userInfo *oauth.WxMiniOauthUserInfo) *model.WxMiniGetUserInfoResp {
	return &model.WxMiniGetUserInfoResp{Nickname: userInfo.Nickname, AvatarUrl: userInfo.AvatarUrl}
}
