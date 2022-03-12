package model

type WxMiniLoginReq struct {
	Code string `json:"code" binding:"required"`
}

type WxMiniLoginResp struct {
	Token string `json:"token"`
}

type WxMiniGetUserInfoReq struct {
	EncryptedData string `json:"encrypted_data" binding:"required"`
	Iv            string `json:"iv" binding:"required"`
}

type WxMiniGetUserInfoResp struct {
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}
