package service

import (
	"context"
	"errors"
	"fmt"
	"ginson/app/model"
	"ginson/pkg/code"
	"ginson/pkg/conf"
	"ginson/pkg/constant"
	"ginson/pkg/log"
	"ginson/pkg/oauth"
)

// 测试内存暂存
var openIdCacheMap = map[string]string{}

var userInfoCacheMap = map[string]*model.WxMiniGetUserInfoResp{}

var wxMiniOauthHandler = oauth.NewWxMiniOauthHandler(conf.AppConf.Ext.WxMiniAppId, conf.AppConf.Ext.WxMiniAppSecret)

type WxMiniService struct {
	wxMiniOauthHandler *oauth.WxMiniOauthHandler
}

var wxMiniService = &WxMiniService{wxMiniOauthHandler: wxMiniOauthHandler}

func GetWxMiniService() *WxMiniService {
	return wxMiniService
}

func (w *WxMiniService) WxMiniLogin(ctx context.Context, req *model.WxMiniLoginReq) (*model.WxMiniLoginResp, code.BizErr) {
	sessionInfo, err := w.wxMiniOauthHandler.CodeToSessionKey(req.Code)
	if err != nil {
		log.Error(fmt.Sprintf("code to session key failed, error: %v", err))
		return nil, code.BizError(err)
	}

	openIdCacheMap[sessionInfo.Openid] = sessionInfo.SessionKey
	token, err := tokenService.createToken(ctx, sessionInfo.Openid, constant.OauthSourceWxMini)
	if err != nil {
		return nil, code.BizError(err)
	}

	return &model.WxMiniLoginResp{Token: token}, nil
}

func (w *WxMiniService) WxMiniGetUserInfo(ctx context.Context, req *model.WxMiniGetUserInfoReq) (*model.WxMiniGetUserInfoResp, code.BizErr) {
	openId, ok := ctx.Value("openId").(string)
	if !ok {
		return nil, code.BizError(errors.New("get open id failed"))
	}
	sessionKey, ok := openIdCacheMap[openId]
	if !ok {
		return nil, code.BizError(errors.New("get session key failed"))
	}
	userInfo, err := w.wxMiniOauthHandler.GetUserInfo(sessionKey, req.EncryptedData, req.Iv)
	if err != nil {
		log.Error(fmt.Sprintf("code to session key failed, error: %v", err))
		return nil, code.BizError(err)
	}

	result := &model.WxMiniGetUserInfoResp{Nickname: userInfo.Nickname, AvatarUrl: userInfo.AvatarUrl}
	userInfoCacheMap[openId] = result
	return result, nil
}
