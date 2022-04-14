package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"ginson/foundation/util"
)

type WxMiniOauthHandler struct {
	*BaseOauthHandler
	BaseOauthConfig
}

type WxMiniSessionKey struct {
	CommonErrResp
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
}

type WxMiniOauthToken struct {
	CommonErrResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type WxMiniOauthUserInfo struct {
	OpenId    string         `json:"openId"`
	UnionId   string         `json:"unionId"`
	NickName  string         `json:"nickName"`
	Language  string         `json:"language"`
	Gender    int            `json:"gender"` // 0未知 1为男性，2为女性
	Province  string         `json:"province"`
	City      string         `json:"city"`
	Country   string         `json:"country"`
	AvatarUrl string         `json:"avatarUrl"`
	Watermark *WatermarkInfo `json:"watermark"`
}

type WxMiniOauthGetPhoneInfoReq struct {
	Code string `json:"code"`
}

type WxMiniOauthPhoneInfo struct {
	PhoneNumber     string         `json:"phoneNumber"`
	PurePhoneNumber string         `json:"purePhoneNumber"`
	CountryCode     int            `json:"countryCode"`
	Watermark       *WatermarkInfo `json:"watermark"`
}

type WxMiniOauthUserPhone struct {
	CommonErrResp
	PhoneInfo *WxMiniOauthPhoneInfo `json:"phone_info"`
}

// NewWxMiniOauthHandler 微信小程序授权工具
func NewWxMiniOauthHandler(appId, appSecret string) *WxMiniOauthHandler {
	return &WxMiniOauthHandler{
		BaseOauthHandler: baseOauthHandler,
		BaseOauthConfig: BaseOauthConfig{
			appId:     appId,
			appSecret: appSecret,
		},
	}
}

// CodeToSessionKey 小程序触发wx.login()获取code，再获取微信小程序授权凭证
func (w *WxMiniOauthHandler) CodeToSessionKey(ctx context.Context, code string) (*WxMiniSessionKey, error) {
	url := w.buildCodeToSessionKeyUrl(code)
	result := &WxMiniSessionKey{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

func (w *WxMiniOauthHandler) buildCodeToSessionKeyUrl(code string) string {
	url := NewUrlHelper(wxMiniOauthCode2TokenUrl).
		AddParam("grant_type", grantTypeAuthorizationCode).
		AddParam("appid", w.appId).
		AddParam("secret", w.appSecret).
		AddParam("js_code", code).
		Build()
	return url
}

// GetUserInfo 小程序wx.getUserProfile()获取加密信息，解密加密串，获取微信小程序授权用户的信息
func (w *WxMiniOauthHandler) GetUserInfo(sessionKey, encryptedData, iv string) (*WxMiniOauthUserInfo, error) {
	decrypt := NewMiniOauthDataDecrypt(sessionKey, encryptedData, iv)
	userInfoJson, err := decrypt.Decrypt()
	if err != nil {
		return nil, err
	}

	result := &WxMiniOauthUserInfo{}
	err = json.Unmarshal([]byte(userInfoJson), &result)
	if err != nil {
		return nil, err
	}

	// 校验水印
	if result.Watermark != nil && !result.Watermark.WatermarkValidate(w.appId) {
		return nil, errors.New("watermark appId invalid")
	}

	return result, nil
}

// GetUserPhone 获取微信小程序授权用户的手机
func (w *WxMiniOauthHandler) GetUserPhone(ctx context.Context, code, accessToken string) (*WxMiniOauthUserPhone, error) {
	url := w.buildUserPhoneUrl(accessToken)
	req := &WxMiniOauthGetPhoneInfoReq{Code: code}
	result := &WxMiniOauthUserPhone{}
	err := util.Post(ctx, url, req, &result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

func (w *WxMiniOauthHandler) buildUserPhoneUrl(accessToken string) string {
	url := NewUrlHelper(wxMiniOauthGetPhoneUrl).
		AddParam("access_token", accessToken).
		Build()
	return url
}

// GetAccessToken 获取微信小程序后台调用accessToken
func (w *WxMiniOauthHandler) GetAccessToken(ctx context.Context) (*WxMiniOauthToken, error) {
	url := w.buildAccessTokenUrl()
	result := &WxMiniOauthToken{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

func (w *WxMiniOauthHandler) buildAccessTokenUrl() string {
	url := NewUrlHelper(wxMiniOauthAccessTokenUrl).
		AddParam("grant_type", grantTypeClientCredential).
		AddParam("appid", w.appId).
		AddParam("secret", w.appSecret).
		Build()
	return url
}
