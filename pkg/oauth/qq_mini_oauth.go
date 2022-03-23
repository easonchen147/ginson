package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"ginson/pkg/utils"
)

type QQMiniOauthHandler struct {
	*BaseOauthHandler
	BaseOauthConfig
}

type QQMiniSessionKey struct {
	CommonErrResp
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
}

type QQMiniOauthUserInfo struct {
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

// NewQQMiniOauthHandler QQ小程序授权工具
func NewQQMiniOauthHandler(appId, appSecret string) *QQMiniOauthHandler {
	return &QQMiniOauthHandler{
		BaseOauthHandler: baseOauthHandler,
		BaseOauthConfig: BaseOauthConfig{
			appId:     appId,
			appSecret: appSecret,
		},
	}
}

// CodeToSessionKey 小程序触发qq.login()获取code，再获取QQ小程序授权凭证
func (w *QQMiniOauthHandler) CodeToSessionKey(ctx context.Context, code string) (*QQMiniSessionKey, error) {
	url := w.buildCodeToSessionKeyUrl(code)
	result := &QQMiniSessionKey{}
	err := utils.Get(ctx, url, &result)
	if err != nil {
		return nil, err
	}

	if result.Errcode != 0 {
		return nil, fmt.Errorf("errCode: %d errMsg: %s", result.Errcode, result.Errmsg)
	}

	return result, nil
}

func (w *QQMiniOauthHandler) buildCodeToSessionKeyUrl(code string) string {
	url := utils.NewUrlHelper(qqMiniOauthCode2TokenUrl).
		AddParam("grant_type", grantTypeAuthorizationCode).
		AddParam("appid", w.appId).
		AddParam("secret", w.appSecret).
		AddParam("js_code", code).
		Build()
	return url
}

// GetUserInfo 小程序qq.getUserInfo()获取加密信息，解密加密串，获取QQ小程序授权用户的信息
func (w *QQMiniOauthHandler) GetUserInfo(sessionKey, encryptedData, iv string) (*QQMiniOauthUserInfo, error) {
	decrypt := NewMiniOauthDataDecrypt(sessionKey, encryptedData, iv)
	userInfoJson, err := decrypt.Decrypt()
	if err != nil {
		return nil, err
	}

	result := &QQMiniOauthUserInfo{}
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
