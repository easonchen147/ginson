package oauth

import (
	"context"
	"fmt"
	"ginson/pkg/util"
)

type QQOauthHandler struct {
	*BaseOauthHandler
	BaseOauthConfig
	redirectUrl string
}

type QQOauthToken struct {
	QQErrResp
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type QQOauthRefreshToken struct {
	QQErrResp
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type QQOauthMe struct {
	QQErrResp
	ClientId string `json:"client_id"`
	Openid   string `json:"openid"`
}

type QQOauthUserInfo struct {
	Ret          int    `json:"ret"`
	Msg          string `json:"msg"`
	Nickname     string `json:"nickname"`
	FigureurlQq1 string `json:"figureurl_qq_1"` // 40*40头像一定有
	FigureurlQq2 string `json:"figureurl_qq_2"` // 100*100头像不一定有
	Gender       string `json:"gender"`         // 男/女/空字符串
}

// NewQQOauthHandler QQ授权工具
func NewQQOauthHandler(appId, appSecret, redirectUrl string) *QQOauthHandler {
	return &QQOauthHandler{
		BaseOauthHandler: baseOauthHandler,
		BaseOauthConfig: BaseOauthConfig{
			appId:     appId,
			appSecret: appSecret,
		},
		redirectUrl: redirectUrl,
	}
}

// GetRedirectUrl 获取授权重定向地址，state可以使用唯一凭证， forMobile = true 表示移动端授权，不传则默认为 PC
func (q *QQOauthHandler) GetRedirectUrl(state string, forMobile bool) (string, error) {
	url := util.NewUrlHelper(qqOauthAuthorizeUrl).
		AddParam("response_type", responseTypeCode).
		AddParam("client_id", q.appId).
		AddParam("redirect_uri", q.redirectUrl).
		AddParam("state", q.getState(state))
	if forMobile {
		url = url.AddParam("display", "mobile") // 显示样式
	}
	return url.Build(), nil
}

// GetAccessToken code换取QQ授权的accessToken
func (q *QQOauthHandler) GetAccessToken(ctx context.Context, code string) (*QQOauthToken, error) {
	url := q.buildAccessTokenUrl(code)
	result := &QQOauthToken{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Error, result.ErrorDescription)
	}

	return result, nil
}

func (q *QQOauthHandler) buildAccessTokenUrl(code string) string {
	url := util.NewUrlHelper(qqOauthAccessTokenUrl).
		AddParam("grant_type", grantTypeAuthorizationCode).
		AddParam("code", code).
		AddParam("client_id", q.appId).
		AddParam("client_secret", q.appSecret).
		AddParam("redirect_uri", q.redirectUrl).
		AddParam("fmt", "json").
		Build()
	return url
}

// GetOpenid 获取QQ授权用户的Openid
func (q *QQOauthHandler) GetOpenid(ctx context.Context, accessToken string) (*QQOauthMe, error) {
	url := q.buildOpenidUrl(accessToken)
	result := &QQOauthMe{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Error, result.ErrorDescription)
	}

	return result, nil
}

func (q *QQOauthHandler) buildOpenidUrl(accessToken string) string {
	url := util.NewUrlHelper(qqOauthMeUrl).
		AddParam("access_token", accessToken).
		AddParam("fmt", "json").
		Build()
	return url
}

// GetUserInfo 获取QQ授权用户的信息
func (q *QQOauthHandler) GetUserInfo(ctx context.Context, openid, accessToken string) (*QQOauthUserInfo, error) {
	url := q.buildUserInfoUrl(openid, accessToken)
	result := &QQOauthUserInfo{}
	err := util.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Ret != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Ret, result.Msg)
	}

	return result, nil
}

func (q *QQOauthHandler) buildUserInfoUrl(openid string, accessToken string) string {
	url := util.NewUrlHelper(qqOauthUserInfoUrl).
		AddParam("access_token", accessToken).
		AddParam("oauth_consumer_key", q.appId).
		AddParam("openid", openid).
		Build()
	return url
}
