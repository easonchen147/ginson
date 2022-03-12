package oauth

import "ginson/pkg/utils"

const (
	// 授权类型
	grantTypeAuthorizationCode = "authorization_code"
	grantTypeRefreshToken      = "refresh_token"

	// 响应类型
	responseTypeCode = "code"
)

// 性别
const (
	Unknown = iota
	Male
	Female
)

// 微信登录授权
const (
	wechatOauthAuthorizeUrl    = "https://open.weixin.qq.com/connect/qrconnect"
	wechatOauthAccessTokenUrl  = "https://api.weixin.qq.com/sns/oauth2/access_token"
	wechatOauthRefreshTokenUrl = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	wechatOauthCheckTokenUrl   = "https://api.weixin.qq.com/sns/auth"
	wechatOauthUserInfoUrl     = "https://api.weixin.qq.com/sns/userinfo"

	wechatOauthScopeLogin = "snsapi_login"
)

// 小程序登录授权
const (
	wxMiniOauthCode2TokenUrl = "https://api.weixin.qq.com/sns/jscode2session"
	wxMiniOauthGetPhoneUrl   = "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
)

// qq授权
const (
	qqOauthAuthorizeUrl   = "https://graph.qq.com/oauth2.0/authorize"
	qqOauthAccessTokenUrl = "https://graph.qq.com/oauth2.0/token"
	qqOauthMeUrl          = "https://graph.qq.com/oauth2.0/me"
	qqOauthUserInfoUrl    = "https://graph.qq.com/user/get_user_info"
)

// WechatCommonErrResp 微信通用错误响应结构
type WechatCommonErrResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// QQErrResp QQ授权错误结构
type QQErrResp struct {
	Error            int    `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// BaseOauthConfig 基础授权配置信息
type BaseOauthConfig struct {
	appId     string
	appSecret string
}

type BaseOauthHandler struct{}

var baseOauthHandler = &BaseOauthHandler{}

func (b *BaseOauthHandler) getState(state string) string {
	if state != "" {
		return state
	}
	return utils.GetUuidV4()
}

// GetGenderByInt 微信、微信小程序的性别都是int，一样的值，不需要转换
func (b *BaseOauthHandler) GetGenderByInt(gender int) int {
	return gender
}

// GetGenderByString qq的性别是中文，需要转换
func (b *BaseOauthHandler) GetGenderByString(gender string) int {
	switch gender {
	case "男":
		return Male
	case "女":
		return Female
	default:
		return Unknown
	}
}
