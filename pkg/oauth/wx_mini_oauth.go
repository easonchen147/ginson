package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"ginson/pkg/utils"
	"github.com/go-resty/resty/v2"
	"github.com/xlstudio/wxbizdatacrypt"
)

type WxMiniOauthHandler struct {
	*BaseOauthHandler
	BaseOauthConfig
}

type WxMiniSessionKey struct {
	WechatCommonErrResp
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
}

type WxMiniOauthToken struct {
	WechatCommonErrResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type WxMiniOauthUserInfo struct {
	Openid    string `json:"openId"`
	Unionid   string `json:"unionid"`
	Nickname  string `json:"nickname"`
	Language  string `json:"language"`
	Gender    int    `json:"gender"` // 0未知 1为男性，2为女性
	Province  string `json:"province"`
	City      string `json:"city"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}

type WxMiniOauthWaterMark struct {
	Timestamp int64  `json:"timestamp"`
	Appid     string `json:"appid"`
}

type WxMiniOauthPhoneInfo struct {
	PhoneNumber     string                `json:"phoneNumber"`
	PurePhoneNumber string                `json:"purePhoneNumber"`
	CountryCode     int                   `json:"countryCode"`
	Watermark       *WxMiniOauthWaterMark `json:"watermark"`
}

type WxMiniOauthUserPhone struct {
	WechatCommonErrResp
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

// GetSessionKey 小程序触发wx.login()获取code，再获取微信小程序授权凭证
func (w *WxMiniOauthHandler) GetSessionKey(code string) (*WxMiniSessionKey, error) {
	url := utils.NewUrlHelper(wxMiniOauthCode2TokenUrl).
		AddParam("grant_type", grantTypeAuthorizationCode).
		AddParam("appid", w.appId).
		AddParam("secret", w.appSecret).
		AddParam("js_code", code).
		Build()

	resp, err := resty.New().R().Post(url)
	if err != nil {
		return nil, err
	}

	result := &WxMiniSessionKey{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

// GetUserInfo 小程序wx.getUserProfile()获取加密信息，解密加密串，获取微信小程序授权用户的信息
func (w *WxMiniOauthHandler) GetUserInfo(sessionKey, encryptedData, iv string) (*WxMiniOauthUserInfo, error) {
	pc := wxbizdatacrypt.WxBizDataCrypt{AppId: w.appId, SessionKey: sessionKey}
	ret, err := pc.Decrypt(encryptedData, iv, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false
	if err != nil {
		return nil, err
	}

	jsonStr, ok := ret.(string)
	if !ok {
		return nil, errors.New("to json string failed")
	}

	result := &WxMiniOauthUserInfo{}
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserPhone 获取微信小程序授权用户的手机
func (w *WxMiniOauthHandler) GetUserPhone(code, accessToken string) (*WxMiniOauthUserPhone, error) {
	url := utils.NewUrlHelper(wxMiniOauthGetPhoneUrl).
		AddParam("access_token", accessToken).
		Build()

	req := map[string]string{"code": code}
	resp, err := resty.New().R().SetBody(req).Post(url)
	if err != nil {
		return nil, err
	}

	result := &WxMiniOauthUserPhone{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}
