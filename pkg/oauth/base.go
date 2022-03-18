package oauth

import "ginson/pkg/utils"

const (
	// 授权类型
	grantTypeAuthorizationCode = "authorization_code"
	grantTypeRefreshToken      = "refresh_token"
	grantTypeClientCredential  = "client_credential"

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
	wechatOauthQrCodeLoginUrl  = "https://open.weixin.qq.com/connect/qrconnect" // 扫码登录，一般用在PC端
	wechatOauthAccessTokenUrl  = "https://api.weixin.qq.com/sns/oauth2/access_token"
	wechatOauthRefreshTokenUrl = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	wechatOauthCheckTokenUrl   = "https://api.weixin.qq.com/sns/auth"
	wechatOauthUserInfoUrl     = "https://api.weixin.qq.com/sns/userinfo"

	wechatOauthScopeLogin = "snsapi_login" // 网页应用使用这个，也就是qrcode扫码登录使用

	wechatOauthAuthorizeUrl = "https://open.weixin.qq.com/connect/oauth2/authorize" // 微信浏览器内打开H5直接授权登录

	// 以下对于直接在微信浏览器打开登录的授权范围
	wechatOauthScopeBase     = "snsapi_base"     //snsapi_base （不弹出授权页面，直接跳转，只能获取用户openid），
	wechatOauthScopeUserInfo = "snsapi_userinfo" //snsapi_userinfo （弹出授权页面，可通过openid拿到昵称、性别、所在地。并且， 即使在未关注的情况下，只要用户授权，也能获取其信息 ）

	wechatOauthWechatRedirect = "#wechat_redirect"
)

// 小程序登录授权
const (
	wxMiniOauthAccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token"
	wxMiniOauthCode2TokenUrl  = "https://api.weixin.qq.com/sns/jscode2session"
	wxMiniOauthGetPhoneUrl    = "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
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
