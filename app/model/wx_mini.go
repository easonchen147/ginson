package model

type WxMiniLoginReq struct {
	Code string `json:"code" binding:"required"`
	// 传了直接返回userInfo
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

type WxMiniLoginResp struct {
	Token    string                 `json:"token"`
	UserInfo *WxMiniGetUserInfoResp `json:"user_info"`
}

type WxMiniGetUserInfoReq struct {
	EncryptedData string `json:"encrypted_data" binding:"required"`
	Iv            string `json:"iv" binding:"required"`
}

type WxMiniGetUserInfoResp struct {
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}
