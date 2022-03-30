package model

type WxMiniLoginReq struct {
	Code string `json:"code" binding:"required"`
	// 传了直接返回userInfo
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

type WxMiniGetUserInfoReq struct {
	EncryptedData string `json:"encrypted_data" binding:"required"`
	Iv            string `json:"iv" binding:"required"`
}

