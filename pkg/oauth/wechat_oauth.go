package oauth

import (
	"encoding/json"
	"fmt"
	"ginson/pkg/utils"
	"github.com/go-resty/resty/v2"
)

type WechatOauthHandler struct {
	*BaseOauthHandler
	BaseOauthConfig
	redirectUrl string
}

type WechatOauthToken struct {
	WechatCommonErrResp
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Openid       string `json:"openid"`
}

type WechatOauthUserInfo struct {
	WechatCommonErrResp
	Openid     string   `json:"openid"`
	Unionid    string   `json:"unionid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"` //  0未知  1为男性，2为女性
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
}

// NewWechatOauthHandler 微信授权工具
func NewWechatOauthHandler(appId, appSecret, redirectUrl string) *WechatOauthHandler {
	return &WechatOauthHandler{
		BaseOauthHandler: baseOauthHandler,
		BaseOauthConfig: BaseOauthConfig{
			appId:     appId,
			appSecret: appSecret,
		},
		redirectUrl: redirectUrl,
	}
}

// GetRedirectUrl 获取微信授权重定向地址，state可以使用唯一凭证
func (w *WechatOauthHandler) GetRedirectUrl(state string) (string, error) {
	url := utils.NewUrlHelper(wechatOauthAuthorizeUrl).
		AddParam("response_type", responseTypeCode).
		AddParam("appid", w.appId).
		AddParam("redirect_uri", w.redirectUrl).
		AddParam("scope", wechatOauthScopeLogin).
		AddParam("state", w.getState(state)).
		Build()
	return url, nil
}

// GetAccessToken code换取微信授权的accessToken
func (w *WechatOauthHandler) GetAccessToken(code string) (*WechatOauthToken, error) {
	url := utils.NewUrlHelper(wechatOauthAccessTokenUrl).
		AddParam("grant_type", grantTypeAuthorizationCode).
		AddParam("code", code).
		AddParam("appid", w.appId).
		AddParam("secret", w.appSecret).
		Build()

	resp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	result := &WechatOauthToken{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

// RefreshToken 刷新accessToken有效期
func (w *WechatOauthHandler) RefreshToken(refreshToken string) (*WechatOauthToken, error) {
	url := utils.NewUrlHelper(wechatOauthRefreshTokenUrl).
		AddParam("grant_type", grantTypeRefreshToken).
		AddParam("refresh_token", refreshToken).
		AddParam("appid", w.appId).
		Build()

	resp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	result := &WechatOauthToken{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

// CheckToken 检查accessToken是否有效
func (w *WechatOauthHandler) CheckToken(accessToken, openId string) error {
	url := utils.NewUrlHelper(wechatOauthCheckTokenUrl).
		AddParam("access_token", accessToken).
		AddParam("openid", openId).
		Build()

	resp, err := resty.New().R().Get(url)
	if err != nil {
		return err
	}

	result := &WechatCommonErrResp{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return err
	}
	if result.Errcode != 0 {
		return fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return nil
}

// GetUserInfo 获取微信授权用户的信息
func (w *WechatOauthHandler) GetUserInfo(openId, accessToken string) (*WechatOauthUserInfo, error) {
	url := utils.NewUrlHelper(wechatOauthUserInfoUrl).
		AddParam("openid", openId).
		AddParam("access_token", accessToken).
		Build()

	resp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	result := &WechatOauthUserInfo{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}
