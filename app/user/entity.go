package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OpenId   string `gorm:"not null; uniqueIndex:open_id_source_uniq,priority:1"` // 第三方登录openId
	Source   string `gorm:"not null; uniqueIndex:open_id_source_uniq,priority:2"` // 哪个平台登录
	Nickname string `gorm:"not null"`
	Avatar   string
	Age      int
	Gender   int // 0=unknown 1=male 2=female
}

type Info struct {
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
