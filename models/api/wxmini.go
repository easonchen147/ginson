package api

type LoginReq struct {
	Code          string `json:"code" binding:"required"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}
