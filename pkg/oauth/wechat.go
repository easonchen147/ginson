package oauth

import (
	"context"
	"fmt"
	"ginson/pkg/util"
)

type WechatOauthHandler struct {
	*BaseOauthHandler
	BaseOauthConfig
	redirectUrl string
}

type WechatOauthToken struct {
	CommonErrResp
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Openid       string `json:"openid"`
}

type WechatOauthUserInfo struct {
	CommonErrResp
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

// GetQrCodeRedirectUrl 获取微信授权重定向地址进行扫码，state可以使用唯一凭证
func (w *WechatOauthHandler) GetQrCodeRedirectUrl(state string) (string, error) {
	url := util.NewUrlHelper(wechatOauthQrCodeLoginUrl).
		AddParam("response_type", responseTypeCode).
		AddParam("appid", w.appId).
		AddParam("redirect_uri", w.redirectUrl).
		AddParam("scope", wechatOauthScopeLogin).
		AddParam("state", w.getState(state)).
		Build()
	return url, nil
}

// GetAuthorizeCodeUrl 微信内部H5登录获取code
func (w *WechatOauthHandler) GetAuthorizeCodeUrl(ctx context.Context, state string) (string, error) {
	url := util.NewUrlHelper(wechatOauthAuthorizeUrl).
		AddParam("response_type", responseTypeCode).
		AddParam("appid", w.appId).
		AddParam("redirect_uri", w.redirectUrl).
		AddParam("scope", wechatOauthScopeUserInfo). // 默认获取用户信息权限
		AddParam("state", w.getState(state)).
		Build()
	return url + wechatOauthWechatRedirect, nil
}

// GetAccessToken code换取微信授权的accessToken
func (w *WechatOauthHandler) GetAccessToken(ctx context.Context, code string) (*WechatOauthToken, error) {
	url := w.buildAccessTokenUrl(code)
	result := &WechatOauthToken{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

func (w *WechatOauthHandler) buildAccessTokenUrl(code string) string {
	url := util.NewUrlHelper(wechatOauthAccessTokenUrl).
		AddParam("grant_type", grantTypeAuthorizationCode).
		AddParam("code", code).
		AddParam("appid", w.appId).
		AddParam("secret", w.appSecret).
		Build()
	return url
}

// RefreshToken 刷新accessToken有效期
func (w *WechatOauthHandler) RefreshToken(ctx context.Context, refreshToken string) (*WechatOauthToken, error) {
	url := w.buildRefreshTokenUrl(refreshToken)
	result := &WechatOauthToken{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

func (w *WechatOauthHandler) buildRefreshTokenUrl(refreshToken string) string {
	url := util.NewUrlHelper(wechatOauthRefreshTokenUrl).
		AddParam("grant_type", grantTypeRefreshToken).
		AddParam("refresh_token", refreshToken).
		AddParam("appid", w.appId).
		Build()
	return url
}

// CheckToken 检查accessToken是否有效
func (w *WechatOauthHandler) CheckToken(ctx context.Context, accessToken, openId string) error {
	url := w.buildCheckTokenUrl(accessToken, openId)
	result := &CommonErrResp{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return err
	}

	if result.Errcode != 0 {
		return fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return nil
}

func (w *WechatOauthHandler) buildCheckTokenUrl(accessToken string, openId string) string {
	url := util.NewUrlHelper(wechatOauthCheckTokenUrl).
		AddParam("access_token", accessToken).
		AddParam("openid", openId).
		Build()
	return url
}

// GetUserInfo 获取微信授权用户的信息
func (w *WechatOauthHandler) GetUserInfo(ctx context.Context, openId, accessToken string) (*WechatOauthUserInfo, error) {
	url := w.buildUserInfoUrl(openId, accessToken)
	result := &WechatOauthUserInfo{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

func (w *WechatOauthHandler) buildUserInfoUrl(openId string, accessToken string) string {
	url := util.NewUrlHelper(wechatOauthUserInfoUrl).
		AddParam("openid", openId).
		AddParam("access_token", accessToken).
		Build()
	return url
}
