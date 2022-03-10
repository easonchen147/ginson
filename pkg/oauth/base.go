package oauth

// 微信登录授权
const (
	wechatOauthAuthorizeUrl    = "https://open.weixin.qq.com/connect/qrconnect"
	wechatOauthAccessTokenUrl  = "https://api.weixin.qq.com/sns/oauth2/access_token"
	wechatOauthRefreshTokenUrl = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	wechatOauthCheckTokenUrl   = "https://api.weixin.qq.com/sns/auth"
	wechatOauthUserInfoUrl     = "https://api.weixin.qq.com/sns/userinfo"
	wechatOauthScopeLogin      = "snsapi_login"
)

// 小程序登录授权
const (
	wxMiniOauthCode2TokenUrl  = "https://api.weixin.qq.com/sns/jscode2session"
	wxMiniOauthAccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token"
	wxMiniOauthGetPhoneUrl    = "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
)

// WechatCommonErrResp 微信通用错误响应结构
type WechatCommonErrResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// BaseOauthConfig 基础授权配置信息
type BaseOauthConfig struct {
	appId     string
	appSecret string
}
