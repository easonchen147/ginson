package api

type UserVO struct {
	UserId   uint   `json:"userId"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
}

type CreateTokenReq struct {
	OpenId   string `json:"openId" binding:"required"`
	Source   string `json:"source" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
}

type TokenResp struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}
