package thirdparty_oauth

import (
	"encoding/json"
	"ginson/pkg/utils"
	"github.com/go-resty/resty/v2"
)

const (
	wechatOauthAuthorizeUrl = "https://open.weixin.qq.com/connect/qrconnect"
	wechatOauthTokenUrl     = "https://api.weixin.qq.com/sns/oauth2/access_token"
	wechatOauthUserInfoUrl  = "https://api.weixin.qq.com/sns/userinfo"
	wechatOauthScopeLogin   = "snsapi_login"
)

type WechatOauthHandler struct {
	clientId     string
	clientSecret string
	redirectUrl  string
}

type WechatOauthToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Openid       string `json:"openid"`
	Unionid      string `json:"unionid"`
}

type WechatOauthUserInfo struct {
	Openid     string   `json:"openid"`
	Unionid    string   `json:"unionid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"` // 1为男性，2为女性
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
}

func NewWechatOauthHandler(clientId, clientSecret, redirectUrl string) *WechatOauthHandler {
	return &WechatOauthHandler{
		clientId:     clientId,
		clientSecret: clientSecret,
		redirectUrl:  redirectUrl,
	}
}

// GetRedirectUrl 获取微信登录重定向地址，state可以使用唯一凭证
func (w *WechatOauthHandler) GetRedirectUrl(state string) (string, error) {
	url := utils.NewUrlHelper(wechatOauthAuthorizeUrl).
		AddParam("response_type", "code").
		AddParam("appid", w.clientId).
		AddParam("redirect_uri", w.redirectUrl).
		AddParam("scope", wechatOauthScopeLogin).
		AddParam("state", state).
		Build()
	return url, nil
}

// GetAccessToken code换去授权的accessToken
func (w *WechatOauthHandler) GetAccessToken(code string) (*WechatOauthToken, error) {
	url := utils.NewUrlHelper(wechatOauthTokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("appid", w.clientId).
		AddParam("secret", w.clientSecret).
		Build()

	resp, err := resty.New().R().Post(url)
	if err != nil {
		return nil, err
	}

	result := &WechatOauthToken{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserInfo 获取授权用户的信息
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
	return result, nil
}
